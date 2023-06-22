package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
