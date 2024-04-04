package usecases

import (
	"notification-service/internal/module/notification/repositories"
)

type usecases struct {
	repo repositories.Repositories
}

type Usecases interface {
}

func New(repo repositories.Repositories) Usecases {
	return &usecases{
		repo: repo,
	}
}
