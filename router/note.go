package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type NoteResource struct {
	Db    *gorm.DB
	Redis *redis.Client
}

func (r *NoteResource) GetNoteList(context *gin.Context) {

}
