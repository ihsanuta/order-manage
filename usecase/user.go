package usecase

import (
	"context"
	"order-manage/models"
	"order-manage/repository"
)

type UserUsecase interface {
	GetProfile(ctx context.Context, userID uint64) (user models.User, err error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetProfile(ctx context.Context, userID uint64) (user models.User, err error) {
	user, err = u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return user, err
	}
	return user, nil
}
