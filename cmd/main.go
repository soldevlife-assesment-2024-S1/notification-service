package main

import (
	"context"
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
	logZap := log_internal.SetupLogger()
	log_internal.Init(logZap)
	logger := log_internal.GetLogger()
	cb := httpclient.InitCircuitBreaker(&cfg.HttpClient, cfg.HttpClient.Type)
	httpClient := httpclient.InitHttpClient(&cfg.HttpClient, cb)
	ctx := context.Background()
	// init message stream
	amqp := messagestream.NewAmpq(&cfg.MessageStream)

	// Init Subscriber
	_, err := amqp.NewSubscriber()
	if err != nil {
		logger.Error(ctx, "Failed to create subscriber", err)
	}

	// Init Publisher
	publisher, err := amqp.NewPublisher()
	if err != nil {
		logger.Error(ctx, "Failed to create publisher", err)
	}

	notificationRepo := repositories.New(logger, httpClient, redis)
	notificationUsecase := usecases.New(notificationRepo)
	middleware := middleware.Middleware{
		Repo: notificationRepo,
	}

	validator := validator.New()
	notificationHandler := handler.NotificationHandler{
		Log:       logger,
		Validator: validator,
		Usecase:   notificationUsecase,
		Publish:   publisher,
	}

	var messageRouters []*message.Router

	// incrementnotificationStock, err := messagestream.NewRouter(publisher, "increment_stock_notification_poisoned", "increment_stock_notification_handler", "increment_stock_notification", subscriber, notificationHandler.IncrementnotificationStock)
	// if err != nil {
	// 	logger.Error(ctx, "Failed to create consume_booking_queue router", err)
	// }

	messageRouters = append(messageRouters)

	serverHttp := http.SetupHttpEngine()

	r := router.Initialize(serverHttp, &notificationHandler, &middleware)

	return r, messageRouters

}
