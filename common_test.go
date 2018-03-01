package mailosaur_go_test

import (
	mailosaur "github.com/mailosaur/mailosaur-go"
	"os"
)

func createTestClient() *mailosaur.MailosaurClient {
	apiKey := os.Getenv("MAILOSAUR_API_KEY")
	if apiKey == "" {
		// TODO panic
		apiKey = "<APIKEY>"
	}

	baseURL := os.Getenv("MAILOSAUR_BASE_URL")
	return mailosaur.NewMailosaurClient(apiKey, baseURL)
}
