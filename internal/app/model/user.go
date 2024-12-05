package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func (u *User) BeforCreate() error {
	if len(u.Password) > 0 {
		enc, err := encriptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}
func encriptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
