package service

import (
	"github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type NoteList interface {
	Create(userId int, list notes.NoteList) (int, error)
	GetAll(userId int) ([]notes.NoteList, error)
	GetById(userId, listId int) (notes.NoteList, error)
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
		NoteList:      NewNotesListService(repos.NoteList),
	}
}
