package usermodel

import "airbnb-golang/pkg/common"

type UserRole int

const (
	RoleGuest UserRole = iota + 1
	RoleHost
	RoleAdmin
)

type User struct {
	common.SQLModel
	Email     string   `json:"email" gorm:"column:email"`
	Password  string   `json:"password" gorm:"column:password"`
	FirstName string   `json:"firstName" gorm:"column:first_name"`
	LastName  string   `json:"lastName" gorm:"column:last_name"`
	Phone     string   `json:"phone" gorm:"column:phone"`
	Role      UserRole `json:"role" gorm:"column:role"`
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

func (u *User) GetUserRole() UserRole {
	return u.Role
}
