package model

import (
	"net/http"
)

type Block struct {

}

func GetBlock(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}
