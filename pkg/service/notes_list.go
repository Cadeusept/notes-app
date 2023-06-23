package service

import (
	"github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/repository"
)

type NotesListService struct {
	repo repository.NoteList
}

func NewNotesListService(repo repository.NoteList) *NotesListService {
	return &NotesListService{repo: repo}
}

func (s *NotesListService) Create(userId int, list notes.NoteList) (int, error) {
	return s.repo.Create(userId, list)
}
