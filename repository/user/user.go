package user

import (
	"context"

	"github.com/imrancluster/th-common-payment/conn"
	"github.com/imrancluster/th-common-payment/models"
	"github.com/imrancluster/th-common-payment/repository"
)

type userRepo struct {
}

// CreateUser create user repo
func (ur *userRepo) CreateUser(ctx context.Context, us *models.User) (*models.User, error) {
	db := conn.PostgresDB()
	err := db.Create(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (ur *userRepo) GetUser(ctx context.Context, id int) (*models.User, error) {
	var userData models.User
	db := conn.PostgresDB()

	if err := db.Find(&userData, id).Error; err != nil {
		return nil, err
	}

	return &userData, nil
}

// NewReader exportable repository
func NewUser() repository.UserRepo {
	return &userRepo{}
}
