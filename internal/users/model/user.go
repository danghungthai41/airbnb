package usermodel

import (
	"airbnb-golang/pkg/common"
	"airbnb-golang/utils"
)

// type UserRole int

// const (
// 	RoleGuest UserRole = iota + 1
// 	RoleHost
// 	RoleAdmin
// )

type User struct {
	common.SQLModel
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"-" gorm:"column:password"`
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
	Phone     string `json:"phone" gorm:"column:phone"`
	Role      string `json:"role" gorm:"column:role"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetUserEmail() string {
	return u.Email
}

func (u *User) GetUserRole() string {
	return u.Role
}

func (u *UserCreate) PrepareForCreation() error {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPassword
	u.Role = "guest"

	return nil
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

type UserCreate struct {
	common.SQLModel
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"-" gorm:"column:role"`
}
