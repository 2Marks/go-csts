package auth

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

func (r *Repository) GetUserDetails(username string) (*types.UserAuthDetails, error) {
	rows, err := r.db.Query(
		"SELECT id, name, username, email, password, role, is_active FROM users WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	}

	var user = new(types.UserAuthDetails)
	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.IsActive,
		)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}
