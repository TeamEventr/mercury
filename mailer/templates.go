package mailer

import (
	"bytes"
	"path/filepath"
	"text/template"
)

type OTPEmailData struct {
	Username string
	OTP      string
}

type ConfirmationEmailData struct {
	Username string
}

func LoadAndRenderTemplate(templateName string, data any) (string, error) {
	templatePath := filepath.Join("templates", templateName)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var rendered bytes.Buffer
	err = tmpl.Execute(&rendered, data)
	if err != nil {
		return "", err
	}

	return rendered.String(), nil
}
