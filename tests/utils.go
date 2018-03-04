package tests

import (
	"fmt"
	"github.com/mailosaur/mailosaur-go/mailosaur"
	"math/rand"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func createTestClient() *mailosaur.MailosaurClient {
	client := mailosaur.NewMailosaurClient(testEnvironment.ApiKey)
	if testEnvironment.BaseUrl != "" {
		client.SetBaseUrl(testEnvironment.BaseUrl)
	}
	return client
}

func generateRandomString() string {
	buf := make([]byte, 10)
	for i := range buf {
		buf[i] = chars[rand.Intn(len(chars))]
	}
	return string(buf)
}

func generateEmailAddress(host, server string) string {
	return fmt.Sprintf("%s.%s@%s", generateRandomString(), server, host)
}
