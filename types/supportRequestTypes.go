package types

import (
	"time"
)

type SupportRequest struct {
	Id          int       `json:"id"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatorId   int       `json:"creatorId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ClosedAt    time.Time `json:"closedAt,omitempty" bson:",omitempty"`
}

type SupportRequestComment struct {
	Id            int    `json:"id"`
	Comment       string `json:"comment"`
	CommenterType string `json:"commenterType"`
}

type CreateSupportRequestDTO struct {
	Subject     string `json:"subject" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type GetAllSupportRequestsDTO struct {
	Page        int    `json:"page" validate:"required"`
	PerPage     int    `json:"perPage" validate:"required"`
	SearchQuery string `json:"searchQuery"`
}

type AddSupportRequestCommentDTO struct {
	Id      int    `json:"id" validate:"required"`
	Comment string `json:"comment" validate:"required"`
}

type GetSupportRequestCommentsDTO struct {
	Id int `json:"id" validate:"required"`
}

type CloseSupportRequestDTO struct {
	Id int `json:"id" validate:"required"`
}

type SupportRequestRepository interface {
	Create(SupportRequest) error
	GetAll(GetAllSupportRequestsDTO) (*[]SupportRequest, error)
	GetById(int) (*SupportRequest, error)
	Close(id int) error
	AddComment(supportRequestId int, comment string, commenterId int) error
	GetAllComments(supportRequestId int) (*[]SupportRequestComment, error)
	HasAgentMadeComment(supportRequestId int) bool
	MarkAsProcessing(id int) error
}

type SupportRequestService interface {
	Create(params CreateSupportRequestDTO, creatorId int) error
	GetAll(GetAllSupportRequestsDTO) (*[]SupportRequest, error)
	Comment(params AddSupportRequestCommentDTO, commenterId int, commenterRole string) error
	GetAllComments(GetSupportRequestCommentsDTO) (*[]SupportRequestComment, error)
	Close(CloseSupportRequestDTO) error
}
