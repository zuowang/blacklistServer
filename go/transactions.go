package model

import (
	"net/http"
)

type Transactions struct {

}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

