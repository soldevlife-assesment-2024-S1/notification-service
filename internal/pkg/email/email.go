package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"notification-service/config"
	"notification-service/internal/module/notification/models/request"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

func InitializeSendEmail(cfg *config.EmailConfig) (d *gomail.Dialer, m *gomail.Message) {
	port, err := strconv.Atoi(cfg.SmtpPort)
	if err != nil {
		panic(err)
	}
	d = &gomail.Dialer{Host: cfg.Server, Port: port, Username: cfg.Username, Password: cfg.Password}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: cfg.SkipSSL}

	m = gomail.NewMessage()

	return

}

func ComposeEmail(m *gomail.Message, data *request.SendEmail, attachmentObjects []io.Reader) *gomail.Message {
	// set Header for email
	to := data.To

	splittedTo := strings.Split(data.To, "|")
	if len(splittedTo) > 1 {
		to = splittedTo[1]
	}

	m.SetHeader("From", data.EmailAddress)
	m.SetHeader("To", to)
	if data.Cc != "" {
		m.SetHeader("Cc", data.Cc)
	}
	if data.Bcc != "" {
		m.SetHeader("Bcc", data.Bcc)
	}
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)
	for i, attachmentObject := range attachmentObjects {
		// objectBuffer := make([]byte, 0)

		objectBuffer := new(bytes.Buffer)
		_, err := objectBuffer.ReadFrom(attachmentObject)

		if err != nil {
			fmt.Println(err.Error())
		}

		objectBytes := objectBuffer.Bytes()
		// attachmentObject.Read(objectBuffer)

		contentType := http.DetectContentType(objectBytes)

		m.Attach(data.Attachments[i], gomail.SetCopyFunc(
			func(w io.Writer) error {
				_, err := w.Write(objectBytes)
				fmt.Println(err)
				return err
			}), gomail.SetHeader(map[string][]string{
			"Content-Type": {
				contentType,
			}}))

	}

	return m
}
