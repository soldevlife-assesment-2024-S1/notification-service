package middleware

import (
	"notification-service/internal/module/notification/repositories"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type Middleware struct {
	Log  *otelzap.Logger
	Repo repositories.Repositories
}
