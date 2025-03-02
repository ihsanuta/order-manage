package usecase

import (
	"context"
	"html"
	"order-manage/models"
	"order-manage/repository"
	"order-manage/utils/token"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *models.User) (err error)
	LoginCheck(ctx context.Context, username string, password string) (tokenAuth string, err error)
}

type authUsecase struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
	}
}

func (a *authUsecase) Register(ctx context.Context, user *models.User) (err error) {
	err = a.BeforeSave(user)
	if err != nil {
		return err
	}

	err = a.userRepository.SaveUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *authUsecase) BeforeSave(user *models.User) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	//remove spaces in username
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil

}

func (a *authUsecase) VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (a *authUsecase) LoginCheck(ctx context.Context, username string, password string) (tokenAuth string, err error) {
	user, err := a.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return tokenAuth, err
	}

	err = a.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return tokenAuth, err
	}

	tokenAuth, err = token.GenerateToken(user.ID)
	if err != nil {
		return tokenAuth, err
	}

	return tokenAuth, nil

}
