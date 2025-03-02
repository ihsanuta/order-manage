package repository

import (
	"context"
	"order-manage/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *models.User) (err error)
	GetUserByUsername(ctx context.Context, username string) (user models.User, err error)
	GetUserByID(ctx context.Context, id uint64) (user models.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) SaveUser(ctx context.Context, user *models.User) (err error) {
	err = u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (user models.User, err error) {
	err = u.db.Model(user).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) GetUserByID(ctx context.Context, id uint64) (user models.User, err error) {
	err = u.db.Model(user).Where("id = ?", id).Take(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
