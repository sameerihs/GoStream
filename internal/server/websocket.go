package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func handleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        fmt.Println("Failed to set WebSocket upgrade:", err)
        return
    }

    go handleConnection(conn)
}

func handleConnection(conn *websocket.Conn) {
    defer conn.Close()

    for {
        _, wavData, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("WebSocket read error:", err)
            break
        }

        go func(wavData []byte) {
            flacData, err := ConvertWAVToFLAC(wavData)
            if err != nil {
                fmt.Println("Conversion error:", err)
                conn.WriteMessage(websocket.TextMessage, []byte("Error in conversion"))
                return
            }

            err = conn.WriteMessage(websocket.BinaryMessage, flacData)
            if err != nil {
                fmt.Println("WebSocket write error:", err)
            }
        }(wavData)
    }
}
