package repository

import (
	"database/sql"
	"mood-api/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Password, user.Role)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	row := r.DB.QueryRow("SELECT id, name, email, password, role FROM users WHERE email = $1", email)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role); err != nil {
		return nil, err
	}
	return user, nil
}
