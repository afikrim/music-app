package usecase

import (
	"github.com/afikrim/music-app/repository"
)

type UserUseCase struct {
	UserRepo repository.UserRepo
}

type UserUseCaseInterface interface {
	FetchWithoutID() (interface{}, error)
}

func (uus *UserUseCase) FetchWithoutID() (interface{}, error) {
	result, err := uus.UserRepo.FetchAll()
	if err != nil {
		return nil, err
	}

	resultsWithoutID := []struct {
		Name string
		Age  int
	}{
		{},
	}

	for _, r := range result {
		resultsWithoutID = append(resultsWithoutID, struct {
			Name string
			Age  int
		}{Name: r.Name, Age: r.Age})
	}

	return resultsWithoutID[1:], nil
}
