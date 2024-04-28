package users

import (
	"database/sql"
	"fmt"

	"github.com/2marks/csts/types"
	"github.com/2marks/csts/utils"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user types.CreateUserDTO) error {
	hashedPassword := utils.GeneratePasswordHash(user.Password)

	_, err := r.db.Exec(
		"INSERT INTO USERS(name, username, email, password, role) VALUES(?,?,?,?,?)",
		user.Name, user.Username, user.Email, hashedPassword, user.Role,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetByEmail(email string) (*types.User, error) {
	rows, err := r.db.Query(
		"SELECT * FROM users WHERE email=?",
		email,
	)
	if err != nil {
		return nil, err
	}

	var user = new(types.User)
	for rows.Next() {
		if user, err = scanUserRows(rows); err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *Repository) GetByUsername(username string) (*types.User, error) {
	rows, err := r.db.Query(
		"SELECT * FROM users where email=?",
		username,
	)
	if err != nil {
		return nil, err
	}

	var user = new(types.User)
	for rows.Next() {
		if user, err = scanUserRows(rows); err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *Repository) GetById(id int) (*types.User, error) {
	fmt.Printf("id = %d \n", id)
	rows, err := r.db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	var user = new(types.User)
	for rows.Next() {
		if user, err = scanUserRows(rows); err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *Repository) GetAll(payload types.GetAllUsersDTO) (*[]types.User, error) {
	fmt.Printf("GetAll : page=%d, perPage=%d \n", payload.Page, payload.PerPage)

	rows, err := r.db.Query(
		"SELECT id,name,username,email,role,is_active,created_at,updated_at FROM users LIMIT ? OFFSET ?",
		payload.PerPage, (payload.Page-1)*payload.PerPage,
	)
	if err != nil {
		return nil, err
	}

	var users = make([]types.User, 0)
	for rows.Next() {
		var user types.User

		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Role,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return &users, err
		}

		if user.Id != 0 {
			users = append(users, user)
		}
	}

	return &users, nil
}

func (r *Repository) Activate(id int) error {
	_, err := r.db.Exec("UPDATE users SET is_active=true WHERE id=?", id)
	return err
}

func (r *Repository) Deactivate(id int) error {
	_, err := r.db.Exec("UPDATE users SET is_active=false WHERE id=?", id)
	return err
}

func scanUserRows(rows *sql.Rows) (*types.User, error) {
	var user = new(types.User)

	err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
