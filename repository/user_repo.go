package repository

import "github.com/afikrim/music-app/model"

type UserRepo struct{}

type UserRepoInterface interface {
	FetchAll() ([]model.User, error)
}

func (hr *UserRepo) FetchAll() ([]model.User, error) {
	return []model.User{
		{"1", "Dave", 30},
		{"2", "Will", 12},
		{"3", "Henry", 11},
	}, nil
}
