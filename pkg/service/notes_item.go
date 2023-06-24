package service

import (
	"github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/repository"
)

type NotesItemService struct {
	repo     repository.NoteItem
	listRepo repository.NoteList
}

func NewNotesItemService(repo repository.NoteItem, listRepo repository.NoteList) *NotesItemService {
	return &NotesItemService{repo: repo, listRepo: listRepo}
}

func (s *NotesItemService) Create(userId, listId int, input notes.NoteItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil { // list doesn't exist or doesn't belong to user
		return 0, err
	}

	return s.repo.Create(listId, input)
}

func (s *NotesItemService) GetAll(userId, listId int) ([]notes.NoteItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *NotesItemService) GetById(userId, itemId int) (notes.NoteItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *NotesItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
