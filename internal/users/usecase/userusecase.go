package userUC

import (
	"airbnb-golang/config"
	usermodel "airbnb-golang/internal/users/model"
	"airbnb-golang/utils"
	"errors"

	"golang.org/x/net/context"
)

type IUserRepo interface {
	Create(context.Context, *usermodel.UserCreate) error
	FindDataWithCondition(context.Context, map[string]any) (*usermodel.User, error)
}

type UserUC struct {
	config   *config.Config
	userRepo IUserRepo
}

func NewUserUC(cfg *config.Config, uc IUserRepo) *UserUC {
	return &UserUC{cfg, uc}
}

// func (u *UserUC) FindDataWithCondition(ctx context.Context, id int) (*usermodel.User, error) {

// }

func (uc *UserUC) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := uc.userRepo.FindDataWithCondition(ctx, map[string]any{"email": data.Email})

	if user != nil {
		return errors.New("email is existed")
	}

	if err := data.PrepareForCreation(); err != nil {
		return err
	}

	if err := uc.userRepo.Create(ctx, data); err != nil {
		return err
	}
	return nil
}

func (uc *UserUC) Login(ctx context.Context, data *usermodel.UserLogin) (*utils.Token, error) {

	user, err := uc.userRepo.FindDataWithCondition(ctx, map[string]any{"email": data.Email})

	if err != nil {
		return nil, errors.New("email or password is invalid")
	}

	if err := utils.ComparePassword(user.Password, data.Password); err != nil {
		return nil, errors.New("email or password is invalid")
	}

	token, err := utils.GenerateJWT(utils.TokenPayload{
		Email: user.Email,
		Role:  user.Role,
	}, uc.config)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return token, nil
}
