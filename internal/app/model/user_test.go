package model_test

import (
	"main/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforCreate())
	assert.NotEmpty(t, u.EncryptedPassword)

}
