package middlewares

import (
	"fmt"
	"net/http"

	"github.com/2marks/csts/utils"
)

func withRoleValidator(handlerFunc http.HandlerFunc, role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedInUser := GetUserFromContext(r.Context())
		if loggedInUser["role"] != role {
			utils.WriteErrorToJson(w, http.StatusUnauthorized, fmt.Errorf("permission denied, user not allowed to access this resource"))
			return
		}

		handlerFunc(w, r)
	}
}

func WithAdminRoleValidator(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return withRoleValidator(handlerFunc, "admin")
}

func WithCustomerRoleValidator(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return withRoleValidator(handlerFunc, "customer")
}

func WithAgentRoleValidator(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return withRoleValidator(handlerFunc, "agent")
}
