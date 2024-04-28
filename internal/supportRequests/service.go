package supportrequests

import (
	"fmt"

	"github.com/2marks/csts/types"
)

type SupportRequestService struct {
	repository types.SupportRequestRepository
}

func NewService(repo types.SupportRequestRepository) *SupportRequestService {
	return &SupportRequestService{repository: repo}
}

func (s *SupportRequestService) Create(params types.CreateSupportRequestDTO, creatorId int) error {
	return s.repository.Create(types.SupportRequest{
		Subject:     params.Subject,
		Description: params.Description,
		Status:      "pending",
		CreatorId:   creatorId,
	})
}

func (s *SupportRequestService) GetAll(params types.GetAllSupportRequestsDTO) (*[]types.SupportRequest, error) {
	supportRequests, err := s.repository.GetAll(params)

	if err != nil {
		fmt.Printf("Error occured while fetching support requests %s", err.Error())
		return nil, fmt.Errorf("error fetching support requests")
	}

	return supportRequests, nil
}

func (s *SupportRequestService) Comment(params types.AddSupportRequestCommentDTO, commenterId int, commenterRole string) error {
	supportRequest, err := s.repository.GetById(params.Id)
	if err != nil {
		return err
	}

	if supportRequest.Status == "closed" {
		return fmt.Errorf("you cannot make comments. support request has already been marked as closed")
	}

	hasAgentMadeComment := s.repository.HasAgentMadeComment(params.Id)
	isCustomer := commenterRole == "customer"
	if !hasAgentMadeComment && isCustomer {
		return fmt.Errorf("access denied. you can only make comment on your support request, after an agent has responded")
	}

	isAgent := commenterRole == "agent"
	if !hasAgentMadeComment && isAgent {
		s.repository.MarkAsProcessing(params.Id)
	}

	return s.repository.AddComment(
		params.Id,
		params.Comment,
		commenterId,
	)
}

func (s *SupportRequestService) GetAllComments(params types.GetSupportRequestCommentsDTO) (*[]types.SupportRequestComment, error) {
	comments, err := s.repository.GetAllComments(params.Id)

	if err != nil {
		fmt.Printf("Error occured while fetching support request comments %s", err.Error())
		return nil, fmt.Errorf("error fetching support request comments")
	}

	return comments, nil
}

func (s *SupportRequestService) Close(params types.CloseSupportRequestDTO) error {
	supportRequest, err := s.repository.GetById(params.Id)
	if err != nil {
		return err
	}

	if supportRequest.Status == "closed" {
		return fmt.Errorf("support request has already been marked as closed")
	}

	return s.repository.Close(params.Id)
}
