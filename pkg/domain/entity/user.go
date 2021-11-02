package entity

type User struct {
	ID   uint64 `json:"-"`
	Name string `json:"name" binding:"required"`
}
