package handler

import (
	"notification-service/internal/module/notification/usecases"
	"notification-service/internal/pkg/log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-playground/validator/v10"
)

type NotificationHandler struct {
	Log       log.Logger
	Validator *validator.Validate
	Usecase   usecases.Usecases
	Publish   message.Publisher
}

func (h *NotificationHandler) NotificationQueue(msg *message.Message) error {
	return nil
}

func (h *NotificationHandler) NotificationInvoice(msg *message.Message) error {
	return nil
}

func (h *NotificationHandler) NotificationPayment(msg *message.Message) error {
	return nil
}
