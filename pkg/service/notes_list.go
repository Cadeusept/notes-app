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

func (s *NotesListService) GetAll(userId int) ([]notes.NoteList, error) {
	return s.repo.GetAll(userId)
}

func (s *NotesListService) GetById(userId, listId int) (notes.NoteList, error) {
	return s.repo.GetById(userId, listId)
}
