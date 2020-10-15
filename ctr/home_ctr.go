package ctr

import (
	"github.com/afikrim/music-app/model"
	"github.com/afikrim/music-app/utils/delivery"
	"net/http"
)

type Home struct {
	Res     delivery.CustomJSON
	Payload model.Payload
}

type HomeInterface interface {
	Home(w http.ResponseWriter, r *http.Request)
}

func (h *Home) Home(w http.ResponseWriter, r *http.Request) {
	h.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, h.Payload.NewPayload("success", "", "welcome to api!"))
}
