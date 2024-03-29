package main

import (
	"fmt"
	"my_todolist_api/src/model"
	"my_todolist_api/src/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	engine := gin.Default()

	server := model.Server{Engine: engine}

	router.UserRouter(server, "/user")

	fmt.Println("Server run on port : ", port)
	server.Engine.Run(":" + port)

}
