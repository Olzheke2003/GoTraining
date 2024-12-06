package store_test

import (
	"main/internal/app/model"
	"main/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, dataBaseURL)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Email: "olzhas@gmail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, dataBaseURL)
	defer teardown("users")
	email := "olzhas@gmail.com"
	_, err := s.User().FindEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "olzhas@gmail.com",
	})
	u, err := s.User().FindEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
