package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

type RoomPostgres struct {
	db *sqlx.DB
}

func NewRoomPostgres(db *sqlx.DB) *RoomPostgres {
	return &RoomPostgres{db: db}
}

func (r *RoomPostgres) Create(userId uint64, room entity.Room) (uint64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id uint64
	createRoomQuery := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", roomsTAble)
	row := tx.QueryRow(createRoomQuery, room.Name)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersRoomsQuery := fmt.Sprintf("INSERT INTO %s (user_id, room_id) VALUES ($1, $2)", usersRoomsTable)
	_, err = tx.Exec(createUsersRoomsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
