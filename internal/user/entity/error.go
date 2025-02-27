package entity

import (
	"errors"
	"go-micro-clean/common"
)

var (
	ErrFirstNameIsEmpty = errors.New("first name can not be blank")
	ErrFirstNameTooLong = errors.New("first name too long, max character is 30")
	ErrLastNameIsEmpty  = errors.New("last name can not be blank")
	ErrLastNameTooLong  = errors.New("last name too long, max character is 30")
	ErrEmailIsNotValid  = errors.New("email is not valid")
	ErrPhoneIsNotValid  = errors.New("phone is not valid")
	ErrGenderIsNotValid = errors.New("gender is not valid")
	ErrRoleIsNotValid   = errors.New("role is not valid")
	ErrStatusIsNotValid = errors.New("status is not valid")
	ErrCannotGetUser    = errors.New("cannot get user info")
	ErrCannotCreateUser = errors.New("cannot create new user")
	ErrCannotGetUsers   = errors.New("cannot get users")

	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
