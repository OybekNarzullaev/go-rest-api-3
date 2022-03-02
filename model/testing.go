package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email: "example@some.com",
		Password: "123456",
	}
}