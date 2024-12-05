package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "qwe@gmail.com",
		Password: "password",
	}
}
