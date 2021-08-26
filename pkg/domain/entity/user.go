package entity

type User struct {
	ID       uint64 `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
