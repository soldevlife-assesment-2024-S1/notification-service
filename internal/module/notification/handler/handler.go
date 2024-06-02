package handler

import (
	"fmt"
	"notification-service/internal/module/notification/models/request"
	"notification-service/internal/module/notification/usecases"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type NotificationHandler struct {
	Log       *otelzap.Logger
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
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to marshal payload %v", err))
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to publish to poison queue %v", err))
			return err
		}

		h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to unmarshal payload %v", err))
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
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to marshal payload %v", err))
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to publish to poison queue %v", err))
			return err
		}

		h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to process notification queue %v", err))
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
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to marshal payload %v", err))
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to publish to poison queue %v", err))
			return err
		}

		h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to unmarshal payload %v", err))
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
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to marshal payload %v", err))
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to publish to poison queue %v", err))
			return err
		}

		h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to process notification invoice %v", err))
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
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to marshal payload %v", err))
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to publish to poison queue %v", err))
			return err
		}

		h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to unmarshal payload %v", err))
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
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to marshal payload %v", err))
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to publish to poison queue %v", err))
			return err
		}

		h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to process notification payment %v", err))
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
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to marshal payload %v", err))
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to publish to poison queue %v", err))
			return err
		}

		h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to unmarshal payload %v", err))
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
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to marshal payload %v", err))
			return err
		}

		if err := h.Publish.Publish("poison_queue", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
			h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to publish to poison queue %v", err))
			return err
		}

		h.Log.Ctx(msg.Context()).Error(fmt.Sprintf("Failed to process notification cancel %v", err))
		return err
	}

	return nil
}
