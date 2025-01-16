package auth

import "strings"

func (lr *LoginRequest) Check() error {
	lr.Email = strings.TrimSpace(lr.Email)

	if !emailIsValid(lr.Email) {
		return ErrEmailIsNotValid
	}

	lr.Password = strings.TrimSpace(lr.Password)

	if err := checkPassword(lr.Password); err != nil {
		return err
	}

	return nil
}

func (rr *RegisterRequest) Check() error {

	rr.Email = strings.TrimSpace(rr.Email)

	if !emailIsValid(rr.Email) {
		return ErrEmailIsNotValid
	}

	rr.Password = strings.TrimSpace(rr.Password)

	if err := checkPassword(rr.Password); err != nil {
		return err
	}

	rr.FirstName = strings.TrimSpace(rr.FirstName)

	if err := checkFirstName(rr.FirstName); err != nil {
		return err
	}

	rr.LastName = strings.TrimSpace(rr.LastName)

	if err := checkLastName(rr.LastName); err != nil {
		return err
	}

	return nil
}
