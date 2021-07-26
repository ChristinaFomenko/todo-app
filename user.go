package todo_app

type User struct {
	Id       int `json:"-"db:"id"`
	Name     int `json:"name" binding:"required"`
	Username int `json:"username" binding:"required"`
	Password int `json:"password" binding:"required"`
}
