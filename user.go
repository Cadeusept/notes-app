package notes

type User struct {
	Id       int    `json:"-"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}
