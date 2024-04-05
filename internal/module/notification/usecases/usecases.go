package usecases

import (
	"notification-service/config"
	"notification-service/internal/module/notification/models/request"
	"notification-service/internal/module/notification/repositories"
	"notification-service/internal/pkg/email"
	texttemplate "notification-service/internal/pkg/helpers/text_template"
)

type usecases struct {
	repo     repositories.Repositories
	cfgEmail *config.EmailConfig
}

// NotificationCancel implements Usecases.
func (u *usecases) NotificationCancel(payload request.NotificationMessage) error {
	templateSendBodyEmail, err := texttemplate.PopulateTemplate(texttemplate.TemplateHtmlMessge, payload)
	if err != nil {
		return err
	}

	specSendEmail := request.SendEmail{
		EmailAddress: payload.EmailRecipient,
		To:           templateSendBodyEmail,
		Subject:      "Cancel Payment Notification",
	}

	dial, message := email.InitializeSendEmail(u.cfgEmail)
	message = email.ComposeEmail(message, &specSendEmail, nil)

	if err := dial.DialAndSend(message); err != nil {
		return err
	}

	// send email
	return nil
}

// NotificationInvoice implements Usecases.
func (u *usecases) NotificationInvoice(payload request.NotificationInvoice) error {
	templateSendBodyEmail, err := texttemplate.PopulateTemplate(texttemplate.TemplateHtmlInvoice, payload)
	if err != nil {
		return err
	}

	specSendEmail := request.SendEmail{
		EmailAddress: payload.EmailRecipient,
		To:           templateSendBodyEmail,
		Subject:      "Invoice Notification",
	}

	dial, message := email.InitializeSendEmail(u.cfgEmail)
	message = email.ComposeEmail(message, &specSendEmail, nil)

	if err := dial.DialAndSend(message); err != nil {
		return err
	}

	// send email
	return nil
}

// NotificationPayment implements Usecases.
func (u *usecases) NotificationPayment(payload request.NotificationPayment) error {
	templateSendBodyEmail, err := texttemplate.PopulateTemplate(texttemplate.TemplateHtmlPayment, payload)
	if err != nil {
		return err
	}

	specSendEmail := request.SendEmail{
		EmailAddress: payload.EmailRecipient,
		To:           templateSendBodyEmail,
		Subject:      "Payment Notification",
	}

	dial, message := email.InitializeSendEmail(u.cfgEmail)
	message = email.ComposeEmail(message, &specSendEmail, nil)

	if err := dial.DialAndSend(message); err != nil {
		return err
	}

	// send email
	return nil
}

// NotificationQueue implements Usecases.
func (u *usecases) NotificationQueue(payload request.NotificationMessage) error {
	templateSendBodyEmail, err := texttemplate.PopulateTemplate(texttemplate.TemplateHtmlMessge, payload)
	if err != nil {
		return err
	}

	specSendEmail := request.SendEmail{
		EmailAddress: payload.EmailRecipient,
		To:           templateSendBodyEmail,
		Subject:      "Queue Notification",
	}

	dial, message := email.InitializeSendEmail(u.cfgEmail)
	message = email.ComposeEmail(message, &specSendEmail, nil)

	if err := dial.DialAndSend(message); err != nil {
		return err
	}

	// send email
	return nil
}

type Usecases interface {
	NotificationQueue(payload request.NotificationMessage) error
	NotificationCancel(payload request.NotificationMessage) error
	NotificationInvoice(payload request.NotificationInvoice) error
	NotificationPayment(payload request.NotificationPayment) error
}

func New(repo repositories.Repositories, cfgEmail *config.EmailConfig) Usecases {
	return &usecases{
		repo:     repo,
		cfgEmail: cfgEmail,
	}
}
