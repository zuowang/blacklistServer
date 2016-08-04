package model

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"GetBlock",
		"GET",
		"/chain/blocks/{Block}",
		GetBlock,
	},

	Route{
		"GetChain",
		"GET",
		"/chain",
		GetChain,
	},

	Route{
		"ChaincodeDeploy",
		"POST",
		"/devops/deploy",
		ChaincodeDeploy,
	},

	Route{
		"ChaincodeInvoke",
		"POST",
		"/devops/invoke",
		ChaincodeInvoke,
	},

	Route{
		"ChaincodeOp",
		"POST",
		"/chaincode",
		ChaincodeOp,
	},

	Route{
		"ChaincodeQuery",
		"POST",
		"/devops/query",
		ChaincodeQuery,
	},

	Route{
		"GetPeers",
		"GET",
		"/network/peers",
		GetPeers,
	},

	Route{
		"DeleteUserRegistration",
		"DELETE",
		"/registrar/{enrollmentID}",
		DeleteUserRegistration,
	},

	Route{
		"GetUserEnrollmentCertificate",
		"GET",
		"/registrar/{enrollmentID}/ecert",
		GetUserEnrollmentCertificate,
	},

	Route{
		"GetUserRegistration",
		"GET",
		"/registrar/{enrollmentID}",
		GetUserRegistration,
	},

	Route{
		"GetUserTransactionCertificate",
		"GET",
		"/registrar/{enrollmentID}/tcert",
		GetUserTransactionCertificate,
	},

	Route{
		"RegisterUser",
		"POST",
		"/registrar",
		RegisterUser,
	},

	Route{
		"GetTransaction",
		"GET",
		"/transactions/{UUID}",
		GetTransaction,
	},

}