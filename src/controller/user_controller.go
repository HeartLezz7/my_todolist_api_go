package controller

import (
	"github.com/gin-gonic/gin"
)

func CheckUser(r *gin.Context) {
	r.JSON(200, "CHECK USER")
}
