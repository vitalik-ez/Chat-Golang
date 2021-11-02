package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RoomPostgres struct {
	db *sqlx.DB
}

func NewRoomPostgres(db *sqlx.DB) *RoomPostgres {
	return &RoomPostgres{db: db}
}

func (r *RoomPostgres) Create(room string) error {
	createRoomQuery := fmt.Sprintf("INSERT INTO %s (name) values ($1) ON CONFLICT (name) DO NOTHING", roomsTable)
	_, err := r.db.Exec(createRoomQuery, room)
	if err != nil {
		return err
	}
	return nil

}
