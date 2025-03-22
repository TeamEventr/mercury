package mailer

import (
	"fmt"

	"github.com/resend/resend-go/v2"
)

type Mailer struct {
	client *resend.Client
	from   string
}

func New(apiKey, fromEmail string) (*Mailer, error) {
	if apiKey == "" || fromEmail == "" {
		return nil, fmt.Errorf("API-Key and From email are MANDATORY.")
	}

	client := resend.NewClient(apiKey)
	return &Mailer{
		client: client,
		from:   fromEmail,
	}, nil
}

func (m *Mailer) SendEmail(to, subject, htmlContent string) error {
	if to == "" || subject == "" || htmlContent == "" {
		return fmt.Errorf("to, subject and html-content are required")
	}

	params := &resend.SendEmailRequest{
		From:    m.from,
		To:      []string{to},
		Subject: subject,
		Html:    htmlContent,
	}

	_, err := m.client.Emails.Send(params)
	if err != nil {
		return err
	}

	// TODO: Email response-id can be logged

	return nil
}

func (m *Mailer) SendOTPEmail(to, username, otp string) error {
	data := OTPEmailData{
		Username: username,
		OTP:      otp,
	}

	htmlContent, err := LoadAndRenderTemplate("otp.html", data)
	if err != nil {
		return err
	}

	subject := `Your One-Time Password for LoopIn`
	return m.SendEmail(to, subject, htmlContent)
}

/*
 This is not being implemented for now due to limits with the email sending
 APIs and the lack of a queueing system. When a queueing system is
 implemented at a later point, then we can expand the mailer to include such
 non-functional mails

func (m *Mailer) SendConfirmationEmail(to, username string) error {
	data := ConfirmationEmailData{
		Username: username,
	}

	htmlContent, err := LoadAndRenderTemplate("confirmation.html", data)
	if err != nil {
		return err
	}

	subject := "Welcome Aboard!"
	return m.SendEmail(to, subject, htmlContent)
}
*/
