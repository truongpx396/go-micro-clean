package entity

import (
	"project/common"
)

type Auth struct {
	common.SQLModel
	UserId     int    `json:"user_id" gorm:"column:user_id;"`
	AuthType   string `json:"auth_type" gorm:"column:auth_type;"`
	Email      string `json:"email" gorm:"column:email;"`
	Salt       string `json:"salt" gorm:"column:salt;"`
	Password   string `json:"password" gorm:"column:password;"`
	FacebookId string `json:"facebook_id" gorm:"column:facebook_id"`
}

func (Auth) TableName() string { return "auths" }

func NewAuthWithEmailPassword(userId int, email, salt, password string) Auth {
	return Auth{
		SQLModel: common.NewSQLModel(),
		UserId:   userId,
		Email:    email,
		Salt:     salt,
		Password: password,
		AuthType: "email_password",
	}
}
