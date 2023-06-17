package notes

type User struct {
	Id       int    `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Username string `json:"username"`
}
