package store

import "main/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) FindEmail(email string) (*model.User, error) {
	return nil, nil
}
