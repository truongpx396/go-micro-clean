package entity

import (
	"encoding/json"
	"go-micro-clean/common"

	"github.com/btcsuite/btcutil/base58"
)

type SimpleUser struct {
	common.SQLModel
	LastName  string        `json:"last_name" gorm:"column:last_name;" db:"last_name"`
	FirstName string        `json:"first_name" gorm:"column:first_name;" db:"first_name"`
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar;" db:"avatar"`
}

func (SimpleUser) TableName() string {
	return User{}.TableName()
}

func NewSimpleUser(id int, firstName, lastName string, avatar *common.Image) SimpleUser {
	return SimpleUser{
		SQLModel:  common.SQLModel{ID: uint(id)},
		LastName:  lastName,
		FirstName: firstName,
		Avatar:    avatar,
	}
}

type User struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"-" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      UserRole      `json:"role" gorm:"column:role;"`
	Status    Status        `json:"-" gorm:"column:status" db:"status"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) GetUserId() int {
	return int(u.ID)
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role.String()
}

func (User) TableName() string {
	return "users"
}

type UserCreation struct {
	common.SQLModel
	Email string `json:"email" gorm:"column:email;"`
	// Password  string        `json:"password" gorm:"column:password;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Role      string `json:"-" gorm:"column:roles;"`
	// Salt      string        `json:"-" gorm:"column:salt;"`
	Status Status        `json:"-" gorm:"column:status" db:"status"`
	Avatar *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func NewUserForCreation(firstName, lastName, email string, avatarData string) UserCreation {
	userAvatar := &common.Image{}

	json.Unmarshal(base58.Decode(avatarData), userAvatar)

	return UserCreation{
		SQLModel:  common.NewSQLModel(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Role:      RoleUser.String(),
		Status:    StatusActive,
		Avatar:    userAvatar,
	}
}

func (u *UserCreation) PrepareForInsert() {
	u.SQLModel = common.NewSQLModel()
	u.Role = RoleUser.String()
	u.Status = StatusActive
}

type UserIdRead struct {
	UserId string `json:"id"`
}

func (UserCreation) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

//type Account struct {
//	AccessToken  *tokenprovider.Token `json:"access_token"`
//	RefreshToken *tokenprovider.Token `json:"refresh_token"`
//}
//
//func NewAccount(at, rt *tokenprovider.Token) *Account {
//	return &Account{
//		AccessToken:  at,
//		RefreshToken: rt,
//	}
//}
