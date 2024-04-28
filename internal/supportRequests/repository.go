package supportrequests

import (
	"database/sql"
	"fmt"

	"github.com/2marks/csts/types"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(supportRequest types.SupportRequest) error {
	_, err := r.db.Exec(
		"INSERT INTO support_requests(subject,description,status,creator_id) VALUES(?,?,?,?)",
		supportRequest.Subject, supportRequest.Description, supportRequest.Status, supportRequest.CreatorId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAll(params types.GetAllSupportRequestsDTO) (*[]types.SupportRequest, error) {
	rows, err := r.db.Query(
		`SELECT id,subject,description,status,creator_id,created_at,updated_at
		FROM support_requests LIMIT ? OFFSET ?`,
		params.Page, (params.Page-1)*params.PerPage,
	)

	if err != nil {
		return nil, err
	}

	var supportRequests = make([]types.SupportRequest, 0)

	for rows.Next() {
		supportRequest := new(types.SupportRequest)
		err := rows.Scan(
			&supportRequest.Id,
			&supportRequest.Subject,
			&supportRequest.Description,
			&supportRequest.Status,
			&supportRequest.CreatorId,
			&supportRequest.CreatedAt,
			&supportRequest.UpdatedAt,
			//&supportRequest.ClosedAt,
		)

		if err != nil {
			return nil, err
		}

		if supportRequest.Id != 0 {
			supportRequests = append(supportRequests, *supportRequest)
		}
	}

	return &supportRequests, nil
}

func (r *Repository) GetById(id int) (*types.SupportRequest, error) {
	rows, err := r.db.Query(
		`SELECT id,subject,description,status,creator_id,created_at,updated_at
		 FROM support_requests WHERE id=?`,
		id,
	)

	if err != nil {
		return nil, err
	}

	var supportRequest = new(types.SupportRequest)
	for rows.Next() {
		err := rows.Scan(
			&supportRequest.Id,
			&supportRequest.Subject,
			&supportRequest.Description,
			&supportRequest.Status,
			&supportRequest.CreatorId,
			&supportRequest.CreatedAt,
			&supportRequest.UpdatedAt,
			//&supportRequest.ClosedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	if supportRequest.Id == 0 {
		return nil, fmt.Errorf("support request not found")
	}

	return supportRequest, nil
}

func (r *Repository) Close(id int) error {
	_, err := r.db.Exec(
		"UPDATE support_requests SET status=?,closed_at=CURRENT_TIMESTAMP(),updated_at=CURRENT_TIMESTAMP() WHERE id=?",
		"closed", id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAllComments(supportRequestId int) (*[]types.SupportRequestComment, error) {
	rows, err := r.db.Query(
		`
			SELECT src.id,comment,u.role commenter_type FROM support_request_comments src
			INNER JOIN users u ON src.commenter_id=u.id
			WHERE src.support_request_id=?
		`,
		supportRequestId,
	)

	if err != nil {
		return nil, err
	}

	var comments = make([]types.SupportRequestComment, 0)

	for rows.Next() {
		comment := new(types.SupportRequestComment)

		err := rows.Scan(
			&comment.Id,
			&comment.Comment,
			&comment.CommenterType,
		)
		if err != nil {
			return nil, err
		}

		if comment.Id != 0 {
			comments = append(comments, *comment)
		}
	}

	return &comments, nil
}

func (r *Repository) AddComment(supportRequestId int, comment string, commenterId int) error {
	_, err := r.db.Exec(
		"INSERT INTO support_request_comments(support_request_id, comment, commenter_id) VALUES(?,?,?)",
		supportRequestId, comment, commenterId,
	)

	return err
}

func (r *Repository) HasAgentMadeComment(supportRequestId int) bool {
	rows, err := r.db.Query(
		`SELECT count(src.id) num FROM support_request_comments  src
		 LEFT JOIN users u ON src.commenter_id=u.id
		 WHERE support_request_id=? AND u.role=?
		`,
		supportRequestId, "agent",
	)

	if err != nil {
		return false
	}

	numOfComments := 0
	for rows.Next() {
		rows.Scan(
			&numOfComments,
		)
	}

	return numOfComments > 0
}

func (r *Repository) MarkAsProcessing(id int) error {
	_, err := r.db.Exec(
		"UPDATE support_requests SET status=?,updated_at=CURRENT_TIMESTAMP() WHERE id=?",
		"processing", id,
	)

	if err != nil {
		return err
	}

	return nil
}
