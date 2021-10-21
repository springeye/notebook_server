package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	conf "notebook/config"
	"notebook/docs"
	static2 "notebook/static"
	"notebook/store"

	"notebook/database"
)
import swaggerfiles "github.com/swaggo/files"
import ginSwagger "github.com/swaggo/gin-swagger"

func init() {

}

func SetupServer() *gin.Engine {
	adminDir := static2.AdminStaticDir
	webDir := static2.WebStaticDir

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
	r.Use(gin.Logger(), gin.Recovery(), configCache())

	//register static web
	r.StaticFS("/admin", http.FS(adminDir))
	r.StaticFS("/web", http.FS(webDir))
	r.GET("/", func(context *gin.Context) {
		context.Redirect(301, "/web")
	})
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//register api route
	api := r.Group("/api")
	{
		user := UserResource{Db: db, Redis: database.Redis}

		api.POST("/user/login", user.Login)
		api.POST("/user/register", user.Register)
	}
	notebook := api.Group("/notebook")
	note := api.Group("/note")
	notebook.Use(AuthRequired())
	{
		bookControl := NotebookResource{
			Db: db, Redis: database.Redis,
		}
		notebook.GET("/list", bookControl.GetNotebookList)
	}
	note.Use(AuthRequired())
	{
		noteControl := NoteResource{
			Db: db, Redis: database.Redis,
		}
		note.GET("/list", noteControl.GetNoteList)
	}
	return r
}
func configCache() gin.HandlerFunc {
	var astore store.IStore
	cacheConf := conf.Conf.Cache
	if cacheConf.Type == conf.Memory {
	} else if cacheConf.Type == conf.Redis {
		astore = &store.RedisAdapter{Client: database.Redis}
	} else if conf.Memory == cacheConf.Type {
		astore = &store.MemAdapter{}
	} else {
		panic("config store.type must \"memory\" or \"redis\"")
	}
	return store.NewStore(astore)
}
func AuthRequired() gin.HandlerFunc {

	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		}
		cache := store.Default(context)
		val := cache.Get(fmt.Sprintf("token:%s", token))

		if val == nil {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		} else if val == nil {
			context.Status(http.StatusUnauthorized)
			context.Abort()
		} else {
			context.Next()
		}
	}
}
