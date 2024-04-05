package handler

import (
	"notification-service/internal/module/notification/models/request"
	"notification-service/internal/module/notification/usecases"
	"notification-service/internal/pkg/log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
)

type NotificationHandler struct {
	Log       log.Logger
	Validator *validator.Validate
	Usecase   usecases.Usecases
	Publish   message.Publisher
}

func (h *NotificationHandler) NotificationQueue(msg *message.Message) error {
	msg.Ack()
	var req request.NotificationMessage

	if err := json.Unmarshal(msg.Payload, &req); err != nil {
		// publish to poison queue

		reqPoisoned := request.PoisonedQueue{
			TopicTarget: "notification_queue",
			ErrorMsg:    err.Error(),
			Payload:     req,
		}

		payload, err := json.Marshal(reqPoisoned)
		if err != nil {
			h.Log.Error(msg.Context(), "Failed to marshal payload", err)
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Error(msg.Context(), "Failed to publish to poison queue", err)
			return err
		}

		h.Log.Error(msg.Context(), "Failed to unmarshal payload", err)
		return err
	}

	if err := h.Usecase.NotificationQueue(req); err != nil {
		// publish to poison queue

		reqPoisoned := request.PoisonedQueue{
			TopicTarget: "notification_queue",
			ErrorMsg:    err.Error(),
			Payload:     req,
		}

		payload, err := json.Marshal(reqPoisoned)
		if err != nil {
			h.Log.Error(msg.Context(), "Failed to marshal payload", err)
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Error(msg.Context(), "Failed to publish to poison queue", err)
			return err
		}

		h.Log.Error(msg.Context(), "Failed to process notification queue", err)
		return err
	}

	return nil

}

func (h *NotificationHandler) NotificationInvoice(msg *message.Message) error {
	msg.Ack()
	var req request.NotificationInvoice

	if err := json.Unmarshal(msg.Payload, &req); err != nil {
		// publish to poison queue

		reqPoisoned := request.PoisonedQueue{
			TopicTarget: "notification_invoice",
			ErrorMsg:    err.Error(),
			Payload:     req,
		}

		payload, err := json.Marshal(reqPoisoned)
		if err != nil {
			h.Log.Error(msg.Context(), "Failed to marshal payload", err)
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Error(msg.Context(), "Failed to publish to poison queue", err)
			return err
		}

		h.Log.Error(msg.Context(), "Failed to unmarshal payload", err)
		return err
	}

	if err := h.Usecase.NotificationInvoice(req); err != nil {
		// publish to poison queue

		reqPoisoned := request.PoisonedQueue{
			TopicTarget: "notification_invoice",
			ErrorMsg:    err.Error(),
			Payload:     req,
		}

		payload, err := json.Marshal(reqPoisoned)
		if err != nil {
			h.Log.Error(msg.Context(), "Failed to marshal payload", err)
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Error(msg.Context(), "Failed to publish to poison queue", err)
			return err
		}

		h.Log.Error(msg.Context(), "Failed to process notification invoice", err)
		return err
	}

	return nil
}

func (h *NotificationHandler) NotificationPayment(msg *message.Message) error {
	msg.Ack()
	var req request.NotificationPayment

	if err := json.Unmarshal(msg.Payload, &req); err != nil {
		// publish to poison queue

		reqPoisoned := request.PoisonedQueue{
			TopicTarget: "notification_payment",
			ErrorMsg:    err.Error(),
			Payload:     req,
		}

		payload, err := json.Marshal(reqPoisoned)
		if err != nil {
			h.Log.Error(msg.Context(), "Failed to marshal payload", err)
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Error(msg.Context(), "Failed to publish to poison queue", err)
			return err
		}

		h.Log.Error(msg.Context(), "Failed to unmarshal payload", err)
		return err
	}

	if err := h.Usecase.NotificationPayment(req); err != nil {
		// publish to poison queue

		reqPoisoned := request.PoisonedQueue{
			TopicTarget: "notification_payment",
			ErrorMsg:    err.Error(),
			Payload:     req,
		}

		payload, err := json.Marshal(reqPoisoned)
		if err != nil {
			h.Log.Error(msg.Context(), "Failed to marshal payload", err)
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Error(msg.Context(), "Failed to publish to poison queue", err)
			return err
		}

		h.Log.Error(msg.Context(), "Failed to process notification payment", err)
		return err
	}

	return nil
}

func (h *NotificationHandler) NotificationCancel(msg *message.Message) error {
	msg.Ack()
	var req request.NotificationMessage

	if err := json.Unmarshal(msg.Payload, &req); err != nil {
		// publish to poison queue

		reqPoisoned := request.PoisonedQueue{
			TopicTarget: "notification_cancel",
			ErrorMsg:    err.Error(),
			Payload:     req,
		}

		payload, err := json.Marshal(reqPoisoned)
		if err != nil {
			h.Log.Error(msg.Context(), "Failed to marshal payload", err)
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Error(msg.Context(), "Failed to publish to poison queue", err)
			return err
		}

		h.Log.Error(msg.Context(), "Failed to unmarshal payload", err)
		return err
	}

	if err := h.Usecase.NotificationCancel(req); err != nil {
		// publish to poison queue

		reqPoisoned := request.PoisonedQueue{
			TopicTarget: "notification_cancel",
			ErrorMsg:    err.Error(),
			Payload:     req,
		}

		payload, err := json.Marshal(reqPoisoned)
		if err != nil {
			h.Log.Error(msg.Context(), "Failed to marshal payload", err)
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Error(msg.Context(), "Failed to publish to poison queue", err)
			return err
		}

		h.Log.Error(msg.Context(), "Failed to process notification cancel", err)
		return err
	}

	return nil
}
