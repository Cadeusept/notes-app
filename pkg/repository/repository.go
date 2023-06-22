package repository

import (
	"github.com/Cadeusept/notes-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
}

type NoteList interface {
}

type NoteItem interface {
}

type Repository struct {
	Authorization
	NoteList
	NoteItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
