package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"net/http"
	db2 "notebook/db"
	"notebook/resources"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)
}

func main() {

	setupServer().Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func setupServer() *gin.Engine {
	_, err := db2.Redis.Ping(db2.RedisContext).Result()
	if err != nil {
		panic(err)
	}
	gin.ForceConsoleColor()
	db := db2.Database
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Static("/admin", "./static/admin")

	api := r.Group("/api")
	{
		user := resources.UserResource{Db: db, Redis: db2.Redis}
		api.POST("/user/login", user.Login)
	}
	notebook := api.Group("/notebook")
	note := api.Group("/note")
	notebook.Use(AuthRequired())
	{
		notebook.GET("/list", func(context *gin.Context) {
			context.Status(http.StatusOK)
		})
	}
	note.Use(AuthRequired())
	{
		note.GET("/list", func(context *gin.Context) {
			context.Status(http.StatusOK)
		})
	}
	return r
}
func AuthRequired() gin.HandlerFunc {
	contextLogger := log.WithFields(log.Fields{})
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		}
		contextLogger.Debug("request header token: ", token)
		val, err := db2.Redis.Get(db2.RedisContext, fmt.Sprintf("token:%s", token)).Bytes()
		if err == redis.Nil {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		} else if err != nil {
			contextLogger.Panic(err)
			context.Status(http.StatusInternalServerError)
			context.Abort()
		} else if len(val) <= 0 {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		} else {
			context.Next()
		}
	}
}
