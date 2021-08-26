package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash) values($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email = $1 and password_hash = $2", usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
