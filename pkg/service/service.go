package service

import (
	"github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GenerateToken(login, password string) (string, error)
}

type NoteList interface {
}

type NoteItem interface {
}

type Service struct {
	Authorization
	NoteList
	NoteItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
