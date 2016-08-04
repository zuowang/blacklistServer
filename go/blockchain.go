package model

import (
	"net/http"
)

type Blockchain struct {

}

func GetChain(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

