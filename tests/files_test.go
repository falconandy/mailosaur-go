package tests

import (
	"fmt"
	"github.com/mailosaur/mailosaur-go/mailosaur"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type FilesScope struct {
	client *mailosaur.MailosaurClient
	server string
	email  *mailosaur.Message
}

var (
	filesScope     *FilesScope
	filesScopeOnce sync.Once
)

func getFilesScope() *FilesScope {
	filesScopeOnce.Do(func() {
		client := createTestClient()
		server := testEnvironment.Server

		client.Messages().DeleteAll(server)

		testEmailAddress := fmt.Sprintf("wait_for_test.%s@%s", server, testEnvironment.SmtpHost)
		err := sendEmail(server, testEmailAddress)
		if err != nil {
			panic(err)
		}

		email, err := client.Messages().WaitFor(server, mailosaur.SearchCriteria{SentTo: testEmailAddress})
		if err != nil {
			panic(err)
		}

		filesScope = &FilesScope{
			client: client,
			server: server,
			email:  email,
		}
	})
	return filesScope
}

func TestFilesGetEmail(t *testing.T) {
	scope := getFilesScope()
	result, err := scope.client.Files().GetEmail(scope.email.ID)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, len(result) > 1)
	assert.Contains(t, string(result), scope.email.Subject)
}

func TestFilesGetAttachment(t *testing.T) {
	scope := getFilesScope()
	attachment := scope.email.Attachments[0]
	result, err := scope.client.Files().GetAttachment(attachment.ID)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, attachment.Length, len(result))
}
