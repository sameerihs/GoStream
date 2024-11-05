package tests

import (
	"GoStream/internal/server"
	"bytes"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func TestWebSocketConversion(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	server.RegisterRoutes(router)

	server := httptest.NewServer(router)
	defer server.Close()

	wsURL := "ws" + server.URL[4:] + "/ws"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("WebSocket connection failed: %v", err)
	}
	defer ws.Close()

	wavData, err := os.ReadFile("tests/testdata/music.wav")
	if err != nil {
		t.Fatalf("Failed to load test WAV file: %v", err)
	}

	err = ws.WriteMessage(websocket.BinaryMessage, wavData)
	if err != nil {
		t.Fatalf("Failed to send WAV data: %v", err)
	}

	ws.SetReadDeadline(time.Now().Add(5 * time.Second))
	_, flacData, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to receive FLAC data: %v", err)
	}


	if len(flacData) == 0 {
		t.Errorf("Received empty FLAC data from WebSocket")
	}

	expectedSignature := []byte{0x66, 0x4C, 0x61, 0x43}
	if !bytes.HasPrefix(flacData, expectedSignature) {
		t.Errorf("Data does not start with expected FLAC signature")
	}
}
