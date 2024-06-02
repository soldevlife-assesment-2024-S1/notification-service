package repositories

import (
	"github.com/redis/go-redis/v9"
	circuit "github.com/rubyist/circuitbreaker"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type repositories struct {
	log         *otelzap.Logger
	httpClient  *circuit.HTTPClient
	redisClient *redis.Client
}

type Repositories interface {
}

func New(log *otelzap.Logger, httpClient *circuit.HTTPClient, redisClient *redis.Client) Repositories {
	return &repositories{
		log:         log,
		httpClient:  httpClient,
		redisClient: redisClient,
	}
}
