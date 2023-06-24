package repository

import (
	"fmt"

	"github.com/Cadeusept/notes-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user notes.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, login, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (notes.User, error) {
	var user notes.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
