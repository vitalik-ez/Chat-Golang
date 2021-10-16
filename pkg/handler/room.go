package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

func (h *Handler) createRoom(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input entity.Room

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Room.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type listOfRoom struct {
	List []string `json:"list" binding:"required"`
}

func (h *Handler) getAllRooms(c *gin.Context) {
	rooms := listOfRoom{}
	/*for room := range db {
		rooms.List = append(rooms.List, room)
	}*/
	c.JSON(http.StatusOK, rooms)
}
