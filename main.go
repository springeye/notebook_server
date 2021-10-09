package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	db2 "notebook/db"
	"notebook/resources"
)

func main() {
	db := db2.Database
	r := gin.Default()
	r.Static("/admin", "./static/admin")

	api := r.Group("/api")
	{
		user := resources.UserResource{Db: db}
		api.POST("/user/login", user.Login)
	}
	notebook := api.Group("/notebook")
	note := api.Group("/note")
	notebook.Use(AuthRequired())
	{
		notebook.POST("", func(context *gin.Context) {

		})
	}
	note.Use(AuthRequired())
	{
		note.POST("", func(context *gin.Context) {

		})
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("x-token")
		uid := context.GetHeader("x-uid")
		if token == "" || uid == "" {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		}
		val, err := db2.Redis.Get(db2.RedisContext, fmt.Sprintf("uid:%s", uid)).Result()
		if err != nil {
			context.Status(http.StatusInternalServerError)
			context.Abort()
		} else if val != token {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		} else {
			context.Next()
		}
	}
}
