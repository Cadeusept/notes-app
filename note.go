package notes

import "errors"

type NoteList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	IdUser int `json:"id_user"`
	IdList int `json:"id_list"`
}

type NoteItem struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Body  string `json:"body" db:"body" binding:"required"`
}

type ListsItem struct {
	IdList int `json:"id_list"`
	IdItem int `json:"id_item"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
