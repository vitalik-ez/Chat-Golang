package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/vitalik-ez/Chat-Golang/pkg/service"
	mock_service "github.com/vitalik-ez/Chat-Golang/pkg/service/mocks"
)

func httpToWs(t *testing.T, url string) string {
	t.Helper()
	return "ws" + url[len("http"):]
}

func WSServer(t *testing.T, hb *hub, h *Handler) (*httptest.Server, *websocket.Conn) {
	t.Helper()

	s := httptest.NewServer(http.HandlerFunc(h.chatRoomWS))

	wsURL := httpToWs(t, s.URL)
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
	type mockBehavior func(s *mock_service.MockRoom, room string)

	testTable := []struct {
		name         string
		command      HubCommand
		mockBehavior mockBehavior
	}{
		{
			name: "join to the room",
			command: HubCommand{
				UserName: "Vitaliy",
				Command:  "join",
				Room:     "kpi",
			},
			mockBehavior: func(s *mock_service.MockRoom, room string) {
				s.EXPECT().Create(room).Return(nil)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			room := mock_service.NewMockRoom(c)
			testCase.mockBehavior(room, testCase.command.Room)

			services := &service.Service{Room: room}

			hub := NewHub()
			h := NewHandler(services, hub)

			s, ws := WSServer(t, hub, h)
			defer s.Close()
			defer ws.Close()
			go hub.Run()

			sendMessage(t, ws, testCase.command)
			time.Sleep(time.Second)

			if _, ok := hub.Rooms[testCase.command.Room]; !ok {
				t.Fatalf("Expected '%+v' in Room map", testCase.command.Room)
			}
		})
	}

}
