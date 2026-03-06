package usecase

import (

	"context"

	"github.com/SauletTheBest/BackendFinancialApplication/internal/repository"
)

type UserUsecase struct {
	userRepo  repository.UserRepository	
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Register(ctx context.Context, email string, password string) error {
	//logic
	return nil
}

func (u *UserUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	//logic

	return "", nil
}