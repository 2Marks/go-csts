package users

import (
	"net/http"

	m "github.com/2marks/csts/middlewares"
	"github.com/2marks/csts/types"
	"github.com/2marks/csts/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	userService types.UserService
}

func NewHandler(service types.UserService) *Handler {
	return &Handler{userService: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", m.WithAdminRoleValidator(h.createUser)).Methods(http.MethodPost)
	router.HandleFunc("/users", m.WithAdminRoleValidator(h.getAllUsers)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", m.WithAdminRoleValidator(h.getUserById)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}/activate", m.WithAdminRoleValidator(h.activateUser)).Methods(http.MethodPut)
	router.HandleFunc("/users/{id:[0-9]+}/deactivate", m.WithAdminRoleValidator(h.deactivateUser)).Methods(http.MethodPut)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.CreateUserDTO)

	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	if err := h.userService.Create(payload); err != nil {
		utils.WriteErrorToJson(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteCreateSuccessResponse(w, "")
}

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.GetAllUsersDTO)
	payload.Page = utils.GetQueryRequestIntVal(r, "page")
	payload.PerPage = utils.GetQueryRequestIntVal(r, "perPage")

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	users, err := h.userService.GetAll(payload)
	if err != nil {
		utils.WriteErrorToJson(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteCreateOkResponse(w, users)
}

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.GetOneUserDTO)
	payload.Id = utils.GetPathRequestIntVal(r, "id")

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	user, err := h.userService.GetById(payload)
	if err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateOkResponse(w, user)
}

func (h *Handler) activateUser(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.ActivateUserDTO)
	payload.Id = utils.GetPathRequestIntVal(r, "id")

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	if err := h.userService.Activate(payload); err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateOkResponse(w, payload)
}

func (h *Handler) deactivateUser(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.DeactivateUserDTO)
	payload.Id = utils.GetPathRequestIntVal(r, "id")

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	if err := h.userService.Deactivate(payload); err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateOkResponse(w, "")
}
