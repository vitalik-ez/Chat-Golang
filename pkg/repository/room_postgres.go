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

	var exist bool
	queryExistRow := fmt.Sprintf("select exists(select 1 from %s where name = '%s')", roomsTAble, room.Name)
	rowExist := r.db.QueryRow(queryExistRow)
	if err := rowExist.Scan(&exist); err != nil {
		return 0, err
	}

	var id uint64

	if exist {
		query := fmt.Sprintf("SELECT id FROM %s WHERE name = '%s'", roomsTAble, room.Name)
		rowExist = r.db.QueryRow(query)
		if err := rowExist.Scan(&id); err != nil {
			return 0, err
		}
		fmt.Println("Room", room.Name, "already exist with id", id)
		return id, nil
	}

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createRoomQuery := fmt.Sprintf("INSERT INTO %s (name, founder_id) VALUES ($1, $2) RETURNING id", roomsTAble)
	row := tx.QueryRow(createRoomQuery, room.Name, userId)
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
