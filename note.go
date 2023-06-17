package notes

type NoteList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	IdUser int `json:"id_user"`
	IdList int `json:"id_list"`
}

type NoteItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ListsItem struct {
	IdList int `json:"id_list"`
	IdItem int `json:"id_item"`
}
