package service

import "github.com/Cadeusept/notes-app/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
