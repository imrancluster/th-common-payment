package repository

import (
	"context"

	"github.com/imrancluster/th-common-payment/models"
)

// UserRepo ..
type UserRepo interface {
	CreateUser(ctx context.Context, us *models.User) (*models.User, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
}
