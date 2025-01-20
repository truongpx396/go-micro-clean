package entity

import (
	"go-micro-clean/common"
	"go-micro-clean/proto/auth"
	"strings"
)

type AuthEmailPassword struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (ad *AuthEmailPassword) Validate() error {
	ad.Email = strings.TrimSpace(ad.Email)

	if !emailIsValid(ad.Email) {
		return ErrEmailIsNotValid
	}

	ad.Password = strings.TrimSpace(ad.Password)

	if err := checkPassword(ad.Password); err != nil {
		return err
	}

	return nil
}

func (a *AuthEmailPassword) FromProto(req *auth.LoginRequest) {
	a.Email = req.Email
	a.Password = req.Password
}

type AuthRegister struct {
	FirstName string        `json:"first_name" form:"first_name"`
	LastName  string        `json:"last_name" form:"last_name"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	AuthEmailPassword
}

func (ar *AuthRegister) Validate() error {
	if err := ar.AuthEmailPassword.Validate(); err != nil {
		return err
	}

	ar.FirstName = strings.TrimSpace(ar.FirstName)

	if err := checkFirstName(ar.FirstName); err != nil {
		return err
	}

	ar.LastName = strings.TrimSpace(ar.LastName)

	if err := checkLastName(ar.LastName); err != nil {
		return err
	}

	return nil
}

func (a *AuthRegister) FromProto(req *auth.RegisterRequest) {
	a.FirstName = req.FirstName
	a.LastName = req.LastName
	a.Email = req.Email
	a.Password = req.Password
	// a.Avatar = req.Avatar
}

func (a *AuthRegister) ToProto() *auth.RegisterResponse {
	return &auth.RegisterResponse{
		Success: true,
	}
}

type Token struct {
	Token string `json:"token"`
	// ExpiredIn in seconds
	ExpiredIn int `json:"expire_in"`
}

type TokenResponse struct {
	AccessToken Token `json:"access_token"`
	// RefreshToken will be used when access token expired
	// to issue new pair access token and refresh token.
	RefreshToken *Token `json:"refresh_token,omitempty"`
}

func (t *TokenResponse) ToProto() *auth.LoginResponse {
	return &auth.LoginResponse{
		AccessToken: t.AccessToken.Token,
		ExpiredIn:   int32(t.AccessToken.ExpiredIn),
	}
}

type SuccessResponse struct {
	Data bool `json:"data"`
}
