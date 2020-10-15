package repository

import "github.com/afikrim/music-app/model"

type UserRepo struct{}

type UserRepoInterface interface {
	FetchAll() ([]model.User, error)
}

func (hr *UserRepo) FetchAll() ([]model.User, error) {
	return []model.User{
		{1, "Dave"},
		{2, "Will"},
		{3, "Henry"},
	}, nil
}
