package delivery

import (
	"encoding/json"
	"net/http"
)

type CustomJSON struct{}

type CustomJSONInterface interface {
	CustomJSONRes(w http.ResponseWriter, key string, value string, httpStatus int, data interface{})
}

func (cj *CustomJSON) CustomJSONRes(w http.ResponseWriter, key string, value string, httpStatus int, data interface{}) {
	w.Header().Set(key, value)
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
