package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/2marks/csts/types"
	"github.com/2marks/csts/utils"
	"github.com/golang-jwt/jwt/v5"
)

type userIdKey string

var userIdWithKey userIdKey = "userId"

type AuthMiddleware struct {
	userRepository types.UserRepository
}

func NewAuthMiddleware(userRepo types.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{userRepository: userRepo}
}

func (a *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := getTokenFromHeader(r)
		if authToken == "" {
			noAuthorizationSentError(w)
			return
		}

		token, err := utils.ValidateAuthToken(authToken)

		if err != nil {
			log.Printf("Error from validating token %s \n", err.Error())
			permissionDeniedError(w)
			return
		}

		if !token.Valid {
			log.Printf("invalid token supplied  \n")
			permissionDeniedError(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userIdStr := claims["userId"].(string)
		userId, _ := strconv.Atoi(userIdStr)

		fmt.Printf("User id from token=%d \n\n", userId)

		user, err := a.userRepository.GetById(userId)
		if err != nil {
			utils.WriteErrorToJson(w, http.StatusUnauthorized, err)
		}

		if !user.IsActive {
			utils.WriteErrorToJson(w, http.StatusUnauthorized, fmt.Errorf("user is not active"))
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, userIdWithKey, map[string]any{"id": user.Id, "role": user.Role})
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetUserFromContext(ctx context.Context) map[string]any {
	user, ok := ctx.Value(userIdWithKey).(map[string]any)

	if !ok {
		return map[string]any{}
	}

	return user
}

func getTokenFromHeader(r *http.Request) string {
	headerToken := r.Header.Get("Authorization")
	if headerToken == "" {
		return ""
	}
	tokenSlice := strings.Split(headerToken, " ")

	return tokenSlice[1]

}

func noAuthorizationSentError(w http.ResponseWriter) {
	utils.WriteErrorToJson(w, http.StatusUnauthorized, fmt.Errorf("no authorization sent"))
}

func permissionDeniedError(w http.ResponseWriter) {
	utils.WriteErrorToJson(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}
