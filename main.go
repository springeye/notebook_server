package main

import (
	"github.com/gin-gonic/gin"
	db2 "notebook/db"
	"notebook/resources"
)

func main() {
	db := db2.Database
	r := gin.Default()
	r.Static("/admin", "./static/admin")
	user := resources.UserResource{Db: db}
	r.GET("/user/login", user.Login)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
