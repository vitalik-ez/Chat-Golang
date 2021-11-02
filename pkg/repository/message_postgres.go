package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (m *MessagePostgres) Create(message entity.Message) error {
	createMessageQuery := fmt.Sprintf("INSERT INTO %s (room_id, username, text_message) SELECT id, $1, $2 FROM %s WHERE name=$3", messagesTable, roomsTable)
	_, err := m.db.Exec(createMessageQuery, message.UserName, message.Text, message.Room)
	if err != nil {
		return err
	}
	return nil

}
