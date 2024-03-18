package email

import (
	"fmt"
	"github/joaltoroc/storicard/config"
	"github/joaltoroc/storicard/internal/transaction"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type (
	email struct {
		cfg config.Config
	}
)

func NewEmail(cfg config.Config) transaction.Notification {
	return &email{cfg}
}

// SendMail implements transaction.Notification.
func (e *email) SendMail(destination string, fileName string, executionID string, body []map[string]string) error {
	m := mail.NewV3Mail()
	m.SetTemplateID(e.cfg.Notification.TemplateID)

	from := mail.NewEmail(e.cfg.Notification.From, e.cfg.Notification.From)
	m.SetFrom(from)

	personalization := mail.NewPersonalization()

	to := mail.NewEmail(destination, destination)
	personalization.AddTos(to)

	personalization.SetDynamicTemplateData("items", body)
	personalization.SetDynamicTemplateData("fileName", fileName)
	personalization.SetDynamicTemplateData("executionID", executionID)

	m.AddPersonalizations(personalization)

	request := sendgrid.GetRequest(e.cfg.Notification.ApiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(response)

	return nil
}
