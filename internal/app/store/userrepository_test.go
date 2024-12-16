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
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, dataBaseURL)
	defer teardown("users")
	email := "olzhas@gmail.com"
	_, err := s.User().FindEmail(email)
	assert.Error(t, err)
	u := model.TestUser(t)
	u.Email = email

	s.User().Create(&model.User(u)
	u, err := s.User().FindEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
