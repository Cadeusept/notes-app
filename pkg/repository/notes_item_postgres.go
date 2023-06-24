package repository

import (
	"fmt"

	"github.com/Cadeusept/notes-app"
	"github.com/jmoiron/sqlx"
)

type NotesItemPostgres struct {
	db *sqlx.DB
}

func NewNotesItemPostgres(db *sqlx.DB) *NotesItemPostgres {
	return &NotesItemPostgres{db: db}
}

func (r *NotesItemPostgres) Create(listId int, item notes.NoteItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, body) values ($1, $2) RETURNING id", itemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Body)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	fmt.Println(itemId)

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (id_list, id_item) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *NotesItemPostgres) GetAll(userId, listId int) ([]notes.NoteItem, error) {
	var items []notes.NoteItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.body FROM %s ti INNER JOIN %s li ON li.id_item=ti.id INNER JOIN %s ul ON ul.id_list=li.id_list WHERE li.id_list=$1 AND ul.id_user=$2",
		itemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *NotesItemPostgres) GetById(userId, itemId int) (notes.NoteItem, error) {
	var item notes.NoteItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.body FROM %s ti INNER JOIN %s li ON li.id_item=ti.id INNER JOIN %s ul ON ul.id_list=li.id_list WHERE ti.id=$1 AND ul.id_user=$2",
		itemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *NotesItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti using %s li, %s ul WHERE ti.id=li.id_item AND li.id_list=ul.id_list AND ul.id_user=$1 AND ti.id=$2",
		itemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)

	return err
}
