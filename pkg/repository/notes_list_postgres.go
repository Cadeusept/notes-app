package repository

import (
	"fmt"
	"strings"

	"github.com/Cadeusept/notes-app"
	"github.com/jmoiron/sqlx"
)

type NotesListPostgres struct {
	db *sqlx.DB
}

func NewNotesListPostgres(db *sqlx.DB) *NotesListPostgres {
	return &NotesListPostgres{db: db}
}

func (r *NotesListPostgres) Create(userId int, list notes.NoteList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var list_id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", listsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&list_id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_list) values ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListsQuery, userId, list_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return list_id, tx.Commit()
}

func (r *NotesListPostgres) GetAll(userId int) ([]notes.NoteList, error) {
	var lists []notes.NoteList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.id_list WHERE ul.id_user = $1",
		listsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *NotesListPostgres) GetById(userId, listId int) (notes.NoteList, error) {
	var list notes.NoteList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.id_list WHERE ul.id_user = $1 AND ul.id_list = $2",
		listsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *NotesListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.id_list AND ul.id_user=$1 AND ul.id_list=$2",
		listsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *NotesListPostgres) Update(userId, listId int, input notes.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.id_list AND ul.id_list=$%d AND ul.id_user=$%d",
		listsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
