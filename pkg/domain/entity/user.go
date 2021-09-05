package entity

type User struct {
	ID       uint64 `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
