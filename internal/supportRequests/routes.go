package supportrequests

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/2marks/csts/middlewares"
	"github.com/2marks/csts/types"
	"github.com/2marks/csts/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	service types.SupportRequestService
}

func NewHandler(service types.SupportRequestService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/support-requests", m.WithCustomerRoleValidator(h.Create)).Methods(http.MethodPost)
	router.HandleFunc("/support-requests", h.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/support-requests/{id:[0-9]+}/comments", h.Comment).Methods(http.MethodPost)
	router.HandleFunc("/support-requests/{id:[0-9]+}/comments", h.GetAllComments).Methods(http.MethodGet)
	router.HandleFunc("/support-requests/{id:[0-9]+}/close", m.WithAgentRoleValidator(h.Close)).Methods(http.MethodPut)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.CreateSupportRequestDTO)

	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErrorToJson(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	loggedInUser := m.GetUserFromContext(r.Context())
	if err := h.service.Create(*payload, loggedInUser["id"].(int)); err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateSuccessResponse(w, "support request created successfully")
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.GetAllSupportRequestsDTO)
	payload.Page = utils.GetQueryRequestIntVal(r, "page")
	payload.PerPage = utils.GetQueryRequestIntVal(r, "perPage")

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	out, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println("Payload is", string(out))

	supportRequests, err := h.service.GetAll(*payload)
	if err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateOkResponse(w, supportRequests)
}

func (h *Handler) Comment(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.AddSupportRequestCommentDTO)

	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErrorToJson(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	loggedInUser := m.GetUserFromContext(r.Context())
	userId := loggedInUser["id"].(int)
	userRole := loggedInUser["role"].(string)
	if err := h.service.Comment(*payload, userId, userRole); err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateSuccessResponse(w, "comment added successfully")
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.GetSupportRequestCommentsDTO)
	payload.Id = utils.GetPathRequestIntVal(r, "id")

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	supportRequestComments, err := h.service.GetAllComments(*payload)
	if err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateOkResponse(w, supportRequestComments)
}

func (h *Handler) Close(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.CloseSupportRequestDTO)
	payload.Id = utils.GetPathRequestIntVal(r, "id")

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, errors)
		return
	}

	if err := h.service.Close(*payload); err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteCreateSuccessResponse(w, "support request closed successfully")
}
