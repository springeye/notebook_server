package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type NotebookResource struct {
	Db    *gorm.DB
	Redis *redis.Client
}

func (r *NotebookResource) GetNotebookList(context *gin.Context) {

}
