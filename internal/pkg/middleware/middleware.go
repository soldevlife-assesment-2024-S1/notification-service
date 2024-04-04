package middleware

import (
	"notification-service/internal/module/notification/repositories"
	log "notification-service/internal/pkg/log"
)

type Middleware struct {
	Log  log.Logger
	Repo repositories.Repositories
}
