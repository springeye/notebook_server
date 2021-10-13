package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	. "notebook/cache"
	conf "notebook/config"
	"notebook/database"
	"notebook/resources"
	"os"
)

//go:embed static
var static embed.FS

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

	setupServer().Run(fmt.Sprintf("0.0.0.0:%d", conf.Config.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupServer() *gin.Engine {
	_, err := database.Redis.Ping(database.RedisContext).Result()
	if err != nil {
		panic(err)
	}
	gin.ForceConsoleColor()
	db := database.Database
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.NoRoute(func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 404,
			"msg":  "not found api route",
		})
	})
	r.Use(gin.Logger(), gin.Recovery())
	//register static web
	adminDir, _ := fs.Sub(static, "static/admin")
	r.StaticFS("/admin", http.FS(adminDir))
	webDir, _ := fs.Sub(static, "static/web")
	r.StaticFS("/web", http.FS(webDir))
	r.GET("/", func(context *gin.Context) {
		context.Redirect(301, "/web")
	})
	//register api route
	api := r.Group("/api")
	{
		user := resources.UserResource{Db: db, Redis: database.Redis}

		api.POST("/user/login", user.Login)
		api.POST("/user/register", user.Register)
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
	ctx := context.Background()
	contextLogger := log.WithFields(log.Fields{})
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		}

		val, err := Cache.Get(ctx, fmt.Sprintf("token:%s", token))

		if err == redis.Nil || err == bigcache.ErrEntryNotFound {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		} else if err != nil {
			contextLogger.Panic(err)
			context.Status(http.StatusInternalServerError)
			context.Abort()
		} else if val == nil {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		} else {
			context.Next()
		}
	}
}
