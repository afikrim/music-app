package ctr

import (
	"github.com/afikrim/music-app/model"
	"github.com/afikrim/music-app/repository"
	"github.com/afikrim/music-app/utils/delivery"
	"net/http"
)

type UserCTR struct {
	Res      delivery.CustomJSON
	Payload  model.Payload
	UserRepo repository.UserRepo
}

type UserCTRInterface interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

func (uc *UserCTR) GetUsers(w http.ResponseWriter, _ *http.Request) {
	result, err := uc.UserRepo.FetchAll()
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusInternalServerError, uc.Payload.NewPayload("error", err.Error(), nil))
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, uc.Payload.NewPayload("success", "", result))
	return
}
