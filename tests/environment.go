package tests

import (
	"os"
	"strconv"
)

type MailosaurTestEnvironment struct {
	ApiKey   string
	BaseUrl  string
	Server   string
	SmtpHost string
	SmtpPort int
}

var (
	testEnvironment MailosaurTestEnvironment
)

func init() {
	apiKey := os.Getenv("MAILOSAUR_API_KEY")
	server := os.Getenv("MAILOSAUR_SERVER")

	if apiKey == "" || server == "" {
		panic("Missing necessary environment variables - refer to README.md")
	}

	baseURL := os.Getenv("MAILOSAUR_BASE_URL")

	host := os.Getenv("MAILOSAUR_SMTP_HOST")
	if host == "" {
		host = "mailosaur.io"
	}

	port, _ := strconv.Atoi(os.Getenv("MAILOSAUR_SMTP_PORT"))
	if port == 0 {
		port = 25
	}

	testEnvironment.ApiKey = apiKey
	testEnvironment.BaseUrl = baseURL
	testEnvironment.Server = server
	testEnvironment.SmtpHost = host
	testEnvironment.SmtpPort = port
}
