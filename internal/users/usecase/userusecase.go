package userUC

import (
	usermodel "airbnb-golang/internal/users/model"

	"golang.org/x/net/context"
)

type IUserRepo interface {
	Create(context.Context, *usermodel.User) error
}

type UserUC struct {
	userUC IUserRepo
}

func NewUserUC(uc IUserRepo) *UserUC {
	return &UserUC{uc}
}

func (uc *UserUC) CreateUser(ctx context.Context, user *usermodel.User) (*usermodel.User, error) {
	if err := uc.userUC.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
