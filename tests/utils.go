package tests

import (
	"github.com/mailosaur/mailosaur-go/mailosaur"
	"math/rand"
	"fmt"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func createTestClient() *mailosaur.MailosaurClient {
	return mailosaur.NewMailosaurClient(testEnvironment.ApiKey, testEnvironment.BaseUrl)
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
