package repositories

import (
	"notification-service/config"
	"notification-service/internal/pkg/log"

	"github.com/redis/go-redis/v9"
	circuit "github.com/rubyist/circuitbreaker"
)

type repositories struct {
	log            log.Logger
	httpClient     *circuit.HTTPClient
	cfgUserService *config.UserService
	redisClient    *redis.Client
}

type Repositories interface {
}

func New(log log.Logger, httpClient *circuit.HTTPClient, redisClient *redis.Client) Repositories {
	return &repositories{
		log:         log,
		httpClient:  httpClient,
		redisClient: redisClient,
	}
}
