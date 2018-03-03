package tests

import (
	"fmt"
	"github.com/mailosaur/mailosaur-go/mailosaur"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
	"time"
)

type EmailsScope struct {
	client *mailosaur.MailosaurClient
	server string
	emails []*mailosaur.MessageSummary
}

var (
	emailsScope     *EmailsScope
	emailsScopeOnce sync.Once
)

func getEmailsScope() *EmailsScope {
	emailsScopeOnce.Do(func() {
		client := createTestClient()
		server := testEnvironment.Server

		client.Messages().DeleteAll(server)
		err := sendEmails(server, 5)
		if err != nil {
			panic(err)
		}

		emails, err := client.Messages().List(server)
		if err != nil {
			panic(err)
		}

		emailsScope = &EmailsScope{
			client: client,
			server: server,
			emails: emails.Items,
		}
	})
	return emailsScope
}

func TestEmailsList(t *testing.T) {
	scope := getEmailsScope()
	assert.Equal(t, 5, len(scope.emails))

	for _, email := range scope.emails {
		validateEmailSummary(t, email)
	}
}

func TestEmailsGet(t *testing.T) {
	scope := getEmailsScope()
	emailToRetrieve := scope.emails[0]
	email, err := scope.client.Messages().Get(emailToRetrieve.ID)
	assert.Nil(t, err)
	validateEmail(t, email)
	validateHeaders(t, email)
}

func TestEmailsGetNotFound(t *testing.T) {
	// Should fail if email is not found
	scope := getEmailsScope()
	_, err := scope.client.Messages().Get(generateRandomString())
	assert.NotNil(t, err)
	_, ok := err.(mailosaur.MailosaurException)
	assert.True(t, ok)
}

func TestEmailsWaitFor(t *testing.T) {
	scope := getEmailsScope()
	testEmailAddress := fmt.Sprintf("wait_for_test.%s@%s", scope.server, testEnvironment.SmtpHost)
	sendEmail(scope.server, testEmailAddress)
	email, err := scope.client.Messages().WaitFor(scope.server, mailosaur.SearchCriteria{
		SentTo: testEmailAddress,
	})
	assert.Nil(t, err)
	validateEmail(t, email)
}

func TestEmailsSearchNoCriteriaError(t *testing.T) {
	scope := getEmailsScope()
	_, err := scope.client.Messages().Search(scope.server, mailosaur.SearchCriteria{})
	assert.NotNil(t, err)
	_, ok := err.(mailosaur.MailosaurException)
	assert.True(t, ok)
}

func TestEmailsSearchBySentTo(t *testing.T) {
	scope := getEmailsScope()
	targetEmail := scope.emails[1]

	result, err := scope.client.Messages().Search(scope.server, mailosaur.SearchCriteria{
		SentTo: targetEmail.To[0].Email,
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, targetEmail.To[0].Email, result.Items[0].To[0].Email)
	assert.Equal(t, targetEmail.Subject, result.Items[0].Subject)
}

func TestEmailsSearchBySentToInvalidEmail(t *testing.T) {
	scope := getEmailsScope()

	criteria := mailosaur.SearchCriteria{
		SentTo: ".not_an_email_address",
	}

	_, err := scope.client.Messages().Search(scope.server, criteria)

	assert.NotNil(t, err)
	_, ok := err.(mailosaur.MailosaurException)
	assert.True(t, ok)
}

func TestEmailsSearchByBody(t *testing.T) {
	scope := getEmailsScope()
	targetEmail := scope.emails[1]
	uniqueString := string([]rune(targetEmail.Subject)[:10])

	result, err := scope.client.Messages().Search(scope.server, mailosaur.SearchCriteria{
		Body: uniqueString + " html",
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, targetEmail.To[0].Email, result.Items[0].To[0].Email)
	assert.Equal(t, targetEmail.Subject, result.Items[0].Subject)
}

func TestEmailsSearchBySubject(t *testing.T) {
	scope := getEmailsScope()
	targetEmail := scope.emails[1]
	uniqueString := string([]rune(targetEmail.Subject)[:10])

	result, err := scope.client.Messages().Search(scope.server, mailosaur.SearchCriteria{
		Subject: uniqueString,
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, targetEmail.To[0].Email, result.Items[0].To[0].Email)
	assert.Equal(t, targetEmail.Subject, result.Items[0].Subject)
}

func TestAnalysisSpam(t *testing.T) {
	scope := getEmailsScope()
	targetId := scope.emails[0].ID
	result, err := scope.client.Analysis().Spam(targetId)
	assert.Nil(t, err)
	for _, rule := range result.SpamFilterResults.SpamAssassin {
		assert.NotEmpty(t, rule.Rule)
		assert.NotEmpty(t, rule.Description)
	}
}

func TestEmailsDelete(t *testing.T) {
	scope := getEmailsScope()
	targetEmailId := scope.emails[4].ID

	err := scope.client.Messages().Delete(targetEmailId)
	assert.Nil(t, err)

	// Attempting to delete again should fail
	err = scope.client.Messages().Delete(targetEmailId)
	assert.NotNil(t, err)
	_, ok := err.(mailosaur.MailosaurException)
	assert.True(t, ok)
}

func validateEmail(t *testing.T, email *mailosaur.Message) {
	validateMetadata(t, email)
	validateAttachmentMetadata(t, email)
	validateHtml(t, email)
	validateText(t, email)
}

func validateEmailSummary(t *testing.T, email *mailosaur.MessageSummary) {
	validateSummaryMetadata(t, email)
	assert.Equal(t, 2, email.Attachments)
}

func validateHtml(t *testing.T, email *mailosaur.Message) {
	// Html.Body
	assert.True(t, strings.HasPrefix(email.Html.Body, "<div dir=\"ltr\">"))

	// Html.Links
	assert.Equal(t, 3, len(email.Html.Links))
	assert.Equal(t, "https://mailosaur.com/", email.Html.Links[0].Href)
	assert.Equal(t, "mailosaur", email.Html.Links[0].Text)
	assert.Equal(t, "https://mailosaur.com/", email.Html.Links[1].Href)
	assert.Empty(t, email.Html.Links[1].Text)
	assert.Equal(t, "http://invalid/", email.Html.Links[2].Href)
	assert.Equal(t, "invalid", email.Html.Links[2].Text)

	// Html.Images
	assert.True(t, strings.HasPrefix(email.Html.Images[1].Src, "cid:"))
	assert.Equal(t, "Inline image 1", email.Html.Images[1].Alt)
}

func validateText(t *testing.T, email *mailosaur.Message) {
	// Text.Body
	assert.True(t, strings.HasPrefix(email.Text.Body, "this is a test"))

	// Text.Links
	assert.Equal(t, 2, len(email.Text.Links))
	assert.Equal(t, "https://mailosaur.com/", email.Text.Links[0].Href)
	assert.Equal(t, email.Text.Links[0].Href, email.Text.Links[0].Text)
	assert.Equal(t, "https://mailosaur.com/", email.Text.Links[1].Href)
	assert.Equal(t, email.Text.Links[1].Href, email.Text.Links[1].Text)
}

func validateHeaders(t *testing.T, email *mailosaur.Message) {
	expectedFromHeader := fmt.Sprintf("\"%s\" <%s>", email.From[0].Name, email.From[0].Email)
	expectedToHeader := fmt.Sprintf("\"%s\" <%s>", email.To[0].Name, email.To[0].Email)
	assert.Equal(t, expectedFromHeader, getEmailHeaderValue(email, "From"))
	assert.Equal(t, expectedToHeader, getEmailHeaderValue(email, "To"))
	assert.Equal(t, email.Subject, getEmailHeaderValue(email, "Subject"))
}

func validateMetadata(t *testing.T, email *mailosaur.Message) {
	validateSummaryMetadata(t, &mailosaur.MessageSummary{
		From:     email.From,
		To:       email.To,
		Cc:       email.Cc,
		Bcc:      email.Bcc,
		Subject:  email.Subject,
		Received: email.Received,
	})
}

func validateSummaryMetadata(t *testing.T, email *mailosaur.MessageSummary) {
	assert.Equal(t, 1, len(email.From))
	assert.Equal(t, 1, len(email.To))
	assert.NotEmpty(t, email.From[0].Email)
	assert.NotEmpty(t, email.From[0].Name)
	assert.NotEmpty(t, email.To[0].Email)
	assert.NotEmpty(t, email.To[0].Name)
	assert.NotEmpty(t, email.Subject)

	assert.Equal(t, time.Now().Format("2006-01-02"), email.Received.Format("2006-01-02"))
}

func validateAttachmentMetadata(t *testing.T, email *mailosaur.Message) {
	assert.Equal(t, 2, len(email.Attachments))

	var file1 = email.Attachments[0]
	assert.NotEmpty(t, file1.ID)
	assert.Equal(t, 82138, file1.Length)
	assert.NotEmpty(t, file1.Url)
	assert.Equal(t, "ii_1435fadb31d523f6", file1.FileName)
	assert.Equal(t, "image/png", file1.ContentType)

	var file2 = email.Attachments[1]
	assert.NotEmpty(t, file2.ID)
	assert.Equal(t, 212080, file2.Length)
	assert.NotEmpty(t, file2.Url)
	assert.Equal(t, "dog.png", file2.FileName)
	assert.Equal(t, "image/png", file2.ContentType)
}

func getEmailHeaderValue(email *mailosaur.Message, field string) string {
	// Fallback casing is used, as header casing is determined by sending server
	headers := email.Metadata.Headers
	fallbackValue := ""
	for _, header := range headers {
		if header.Field == field {
			return header.Value
		}
		if strings.ToLower(header.Field) == strings.ToLower(field) {
			fallbackValue = header.Value
		}
	}
	return fallbackValue
}
