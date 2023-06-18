package userrepo

import (
	usermodel "airbnb-golang/internal/users/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *userRepo {
	return &userRepo{DB}
}

func (r *userRepo) Create(ctx context.Context, user *usermodel.User) error {
	db := r.DB.Begin()
	if err := db.Table(usermodel.User{}.TableName()).Create(user).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}

func (r *userRepo) FindDataWithCondition(ctx context.Context, condition map[string]any) (*usermodel.User, error) {

	db := r.DB.Begin()
	var user usermodel.User
	if err := db.Table(usermodel.User{}.TableName()).Where(condition).First(&user).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, err
	}

	return &user, nil
}
