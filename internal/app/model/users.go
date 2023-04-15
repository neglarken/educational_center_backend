package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id                  int    "json:\"id\""
	Login               string "json:\"login\"" // unique
	UnencryptedPassword string "json:\"password,omitempty\""
	Password            string "json:\"-\""
	FirstName           string "json:\"first_name\""
	LastName            string "json:\"last_name\""
	Surname             string "json:\"surname,omitempty\""      // nullable
	PhoneNumber         string "json:\"phone_number,omitempty\"" // nullable
}

func (u *Users) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.UnencryptedPassword, validation.By(requiredIf(u.Password == "")), validation.Length(6, 100)),
		validation.Field(&u.PhoneNumber, validation.By(requiredIf(u.PhoneNumber != "")), validation.Match(regexp.MustCompile("^(\\+7|8)[0-9]{10}$"))),
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

func (u *Users) Sanitize() {
	u.Password = ""
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
