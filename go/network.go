package model

import (
	"net/http"
)

type Network struct {

}

func GetPeers(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

