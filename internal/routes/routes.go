package routes

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[string]map[*websocket.Conn]bool)

// /interview/gtahefj

func wsHandler(c *gin.Context){
	// return hello response
	fmt.Println("hello")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
	defer conn.Close()

	id := c.Param("id")

	if _, ok := clients[id]; !ok {
		clients[id] = make(map[*websocket.Conn]bool)
	}

	clients[id][conn] = true
	for {
    messageType, p, err := conn.ReadMessage()
    if err != nil {
			delete(clients[id], conn)
            conn.Close()
            break
        }


		for client := range clients[id] {
			// Don't send message to same connection
			if client == conn {
				continue
			}
			err = client.WriteMessage(messageType, p)
			if err != nil {
				delete(clients[id], client)
				client.Close()
			}
		}
	}
}


func RegisterRoutes(router *gin.Engine){
	apiGroup := router.Group("/api/")
	{
		userGroup := apiGroup.Group("/user")
		{
			userGroup.GET("/me", func(c *gin.Context){
				// return hello response
				c.JSON(http.StatusOK, gin.H{
					"message": "hello",
				})				
			})
		}

		wsGroup := apiGroup.Group("/ws")
		{
			wsGroup.GET("/start:id", wsHandler)
		}
	}
}