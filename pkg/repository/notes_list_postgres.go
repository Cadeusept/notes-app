package repository

import (
	"fmt"

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
