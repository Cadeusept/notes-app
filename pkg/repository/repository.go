package repository

import (
	"github.com/Cadeusept/notes-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GetUser(login, password string) (notes.User, error)
}

type NoteList interface {
	Create(userId int, list notes.NoteList) (int, error)
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
		NoteList:      NewNotesListPostgres(db),
	}
}
