package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/vitalik-ez/Chat-Golang/pkg/service"
	mock_service "github.com/vitalik-ez/Chat-Golang/pkg/service/mocks"
)

func httpToWs(t *testing.T, url string) string {
	t.Helper()
	return "ws" + url[len("http"):]
}

func WSServer(t *testing.T, h http.Handler) (*httptest.Server, *websocket.Conn) {
	t.Helper()

	s := httptest.NewServer(h)
	wsURL := httpToWs(t, s.URL)
	fmt.Println(wsURL)
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	return s, ws
}


func sendMessage(t *testing.T, ws *websocket.Conn, command HubCommand) {
	t.Helper()

	if err := ws.WriteJSON(command); err != nil {
		t.Fatalf("%v", err)
	}
}


func TestHandler_chatRoomWS(t *testing.T) {
	type mockBehavior func(s *mock_service)
	hub := NewHub()
	h = Handler{services: &service.Service{}}
	s, ws := WSServer(t, h.chatRoomWS(hub))
	defer s.Close()
	defer ws.Close()

	testTable := []struct{
		name string
		command HubCommand
		mockBehavior mockBehavior
	} {
		{
			name: "join to the room",
			command: &HubCommand{
				UserName: "Vitaliy",
				Command: "join",
				Room: "kpi",
			},
			mockBehavior: func(s *mock_service.)
		},
	}

	for _, testCase := range testTable {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

	
}
