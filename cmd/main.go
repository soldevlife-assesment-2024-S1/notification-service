package main

import (
	"context"
	"fmt"
	"log"
	"notification-service/config"
	"notification-service/internal/module/notification/handler"
	"notification-service/internal/module/notification/repositories"
	"notification-service/internal/module/notification/usecases"
	"notification-service/internal/pkg/http"
	"notification-service/internal/pkg/httpclient"
	log_internal "notification-service/internal/pkg/log"
	"notification-service/internal/pkg/messagestream"
	"notification-service/internal/pkg/middleware"
	"notification-service/internal/pkg/observability"
	"notification-service/internal/pkg/redis"
	router "notification-service/internal/route"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.InitConfig()

	app, messageRouters := initService(cfg)

	for _, router := range messageRouters {
		ctx := context.Background()
		go func(router *message.Router) {
			err := router.Run(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}(router)
	}

	// start http server
	http.StartHttpServer(app, cfg.HttpServer.Port)
}

func initService(cfg *config.Config) (*fiber.App, []*message.Router) {
	redis := redis.SetupClient(&cfg.Redis)
	logZap := log_internal.Setup()
	cb := httpclient.InitCircuitBreaker(&cfg.HttpClient, cfg.HttpClient.Type)
	httpClient := httpclient.InitHttpClient(&cfg.HttpClient, cb)
	ctx := context.Background()
	// init message stream
	amqp := messagestream.NewAmpq(&cfg.MessageStream)

	// Init Subscriber
	subscriber, err := amqp.NewSubscriber()
	if err != nil {
		logZap.Ctx(ctx).Error(fmt.Sprintf("Failed to create subscriber: %v", err))
	}

	// Init Publisher
	publisher, err := amqp.NewPublisher()
	if err != nil {
		logZap.Ctx(ctx).Error(fmt.Sprintf("Failed to create publisher: %v", err))
	}

	notificationRepo := repositories.New(logZap, httpClient, redis)
	notificationUsecase := usecases.New(notificationRepo, &cfg.Email)
	middleware := middleware.Middleware{
		Repo: notificationRepo,
	}

	validator := validator.New()
	notificationHandler := handler.NotificationHandler{
		Log:       logZap,
		Validator: validator,
		Usecase:   notificationUsecase,
		Publish:   publisher,
	}

	var messageRouters []*message.Router

	notificationQueue, err := messagestream.NewRouter(publisher, "notification_queue_poisoned", "notification_handler", "notification", subscriber, notificationHandler.NotificationQueue)

	if err != nil {
		logZap.Ctx(ctx).Error(fmt.Sprintf("Failed to create consume_booking_queue router: %v", err))
	}

	notificationCancel, err := messagestream.NewRouter(publisher, "notification_cancel_poisoned", "notification_cancel_handler", "notification_cancel", subscriber, notificationHandler.NotificationCancel)

	if err != nil {
		logZap.Ctx(ctx).Error(fmt.Sprintf("Failed to create consume_booking_queue router: %v", err))
	}

	notificationInvoice, err := messagestream.NewRouter(publisher, "notification_invoice_poisoned", "notification_invoice_handler", "notification_invoice", subscriber, notificationHandler.NotificationInvoice)
	if err != nil {
		logZap.Ctx(ctx).Error(fmt.Sprintf("Failed to create consume_booking_queue router: %v", err))
	}

	notificationPayment, err := messagestream.NewRouter(publisher, "notification_payment_poisoned", "notification_payment_handler", "notification_payment", subscriber, notificationHandler.NotificationPayment)

	if err != nil {
		logZap.Ctx(ctx).Error(fmt.Sprintf("Failed to create consume_booking_queue router: %v", err))
	}

	messageRouters = append(messageRouters, notificationInvoice, notificationPayment, notificationQueue, notificationCancel)

	serverHttp := http.SetupHttpEngine()
	conn, serviceName, err := observability.InitConn(cfg)
	if err != nil {
		logZap.Ctx(ctx).Fatal(fmt.Sprintf("Failed to create gRPC connection to collector: %v", err))
	}
	// setup log
	observability.InitLogOtel(cfg, serviceName)
	// setup tracer
	observability.InitTracer(conn, serviceName)

	// setup metric
	_, err = observability.InitMeterProvider(conn, serviceName)
	if err != nil {
		logZap.Ctx(ctx).Fatal(fmt.Sprintf("Failed to create meter provider: %v", err))
	}

	r := router.Initialize(serverHttp, &notificationHandler, &middleware)

	return r, messageRouters

}
