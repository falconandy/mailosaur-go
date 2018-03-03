package tests

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

var (
	senderHtml string
	senderText string
)

func init() {
	data, err := ioutil.ReadFile(filepath.Join("resources", "testEmail.html"))
	if err != nil {
		panic(err)
	}
	senderHtml = string(data)

	data, err = ioutil.ReadFile(filepath.Join("resources", "testEmail.txt"))
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
	m.Embed(filepath.Join("resources", "cat.png"), gomail.Rename("ii_1435fadb31d523f6"), gomail.SetHeader(map[string][]string{"Content-Type": {"image/png"}}))
	m.Attach(filepath.Join("resources", "dog.png"), gomail.Rename("dog.png"))

	d := gomail.NewDialer(testEnvironment.SmtpHost, testEnvironment.SmtpPort, "", "")
	return d.DialAndSend(m)
}
