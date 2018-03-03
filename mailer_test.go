package mailosaur_go_test

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

var (
	senderHtml string
	senderText string
)

func init() {
	data, err := ioutil.ReadFile(filepath.Join("test_resources", "testEmail.html"))
	if err != nil {
		panic(err)
	}
	senderHtml = string(data)

	data, err = ioutil.ReadFile(filepath.Join("test_resources", "testEmail.txt"))
	if err != nil {
		panic(err)
	}
	senderText = string(data)
}

func sendEmails(server string, quantity int) error {
	for i := 0; i < quantity; i++ {
		err := sendEmail(server, "")
		if err != nil {
			return err
		}
	}

	// Wait to ensure email has arrived
	time.Sleep(2000 * time.Millisecond)
	return nil
}

func sendEmail(server string, sendToAddress string) error {
	randomString := generateRandomString()
	randomToAddress := sendToAddress
	if randomToAddress == "" {
		randomToAddress = generateEmailAddress(testEnvironment.SmtpHost, server)
	}

	m := gomail.NewMessage()
	m.SetHeader("Subject", randomString+" subject")
	m.SetAddressHeader("From", randomString+"@test.com", fmt.Sprintf("%s %s", randomString, randomString)) //TODO
	m.SetAddressHeader("To", randomToAddress, fmt.Sprintf("%s %s", randomString, randomString))
	m.SetBody("text/plain", strings.Replace(senderText, "REPLACED_DURING_TEST", randomString, -1))
	m.AddAlternative("text/html", strings.Replace(senderHtml, "REPLACED_DURING_TEST", randomString, -1))
	m.Embed(filepath.Join("test_resources", "cat.png"), gomail.Rename("ii_1435fadb31d523f6"), gomail.SetHeader(map[string][]string{"Content-Type": {"image/png"}}))
	m.Attach(filepath.Join("test_resources", "dog.png"), gomail.Rename("dog.png"))

	d := gomail.NewDialer(testEnvironment.SmtpHost, testEnvironment.SmtpPort, "", "")
	return d.DialAndSend(m)
}

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
