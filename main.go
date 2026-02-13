package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	routes "code-hire-pro/internal/routes"
)

func main(){
	// Create a gin route
	r := gin.Default()
	// Import routes from routes file
	routes.RegisterRoutes(r)
	// start the gin server

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()

	fmt.Println("Server is running on port 8080")

}