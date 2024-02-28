package controller

import (
	"context"
	"fmt"
	"my_todolist_api/src/database"
	"my_todolist_api/src/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FindAllUser(r *gin.Context) {
	ctx, close := context.WithTimeout(context.Background(), 5*time.Second)
	mongo := database.MongoDB().Collection(model.Collection.User)

	defer close()

	var filter bson.M
	var result []bson.M

	if jsonErr := r.BindJSON(&filter); jsonErr != nil {
		fmt.Println(jsonErr, "CHECK ERROR")
		return
	}

	value, err := mongo.Find(ctx, filter)
	if err != nil {
		fmt.Println(err, "mongo error")
		r.JSON(500, err)
		return
	}

	if cursorErr := value.All(ctx, &result); cursorErr != nil {
		fmt.Println(cursorErr)
		r.JSON(500, err)
		return
	}
	r.JSON(200, result)

}

func FindUser(r *gin.Context) {
	ctx, close := context.WithTimeout(context.Background(), 5*time.Second)
	mongo := database.MongoDB().Collection(model.Collection.User)

	defer close()

	var filter bson.M
	var user = model.User{}

	if jsonErr := r.BindJSON(&filter); jsonErr != nil {
		fmt.Println(jsonErr, "CHECK ERROR")
	}

	if err := mongo.FindOne(ctx, filter).Decode(&user); err != nil {
		fmt.Println(err, "mongo error")
		r.JSON(500, err)
		return
	}

	r.JSON(200, user)

}
