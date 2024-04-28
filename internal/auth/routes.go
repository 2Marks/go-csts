package auth

import (
	"net/http"

	"github.com/2marks/csts/types"
	"github.com/2marks/csts/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	authService types.AuthService
}

func NewHandler(service types.AuthService) *Handler {
	return &Handler{authService: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/login", h.login).Methods(http.MethodPost)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.LoginDTO)

	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	loginResponse, err := h.authService.Login(*payload)
	if err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateOkResponse(w, *loginResponse)
}
