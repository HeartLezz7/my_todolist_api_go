package router

import (
	"my_todolist_api/src/controller"
	"my_todolist_api/src/model"
)

func UserRouter(server model.Server, path string) {
	router := server.Engine.Group(path)
	router.POST("", controller.FindAllUser)
	router.POST("/one", controller.FindUser)
}
