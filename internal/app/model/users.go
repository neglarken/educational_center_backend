package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id                  int
	Login               string // unique
	UnencryptedPassword string
	Password            string
	FirstName           string
	LastName            string
	Surname             string // nullable
	PhoneNumber         string // nullable
}

func (u *Users) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.UnencryptedPassword, validation.By(requiredIf(u.Password == "")), validation.Length(6, 100)),
		validation.Field(&u.PhoneNumber, validation.By(requiredIf(u.PhoneNumber == "")), validation.Match(regexp.MustCompile("^(\\+7|8)[0-9]{10}$"))),
	)
}

func (u *Users) BeforeCreate() error {
	if len(u.UnencryptedPassword) > 0 {
		enc, err := encryptString(u.UnencryptedPassword)
		if err != nil {
			return err
		}
		u.Password = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
