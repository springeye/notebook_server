package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"notebook/cache"
	"notebook/model"
)

type NotebookResource struct {
	Db    *gorm.DB
	Redis *redis.Client
}

func (r *NotebookResource) GetNotebookList(c *gin.Context) {
	token := c.GetHeader("Authorization")
	ctx := context.Background()
	rr, err := cache.Cache.Get(ctx, fmt.Sprintf("token:%s", token))
	if err != nil {
		sendError(c, 9999, err.Error())
	}
	var result model.User
	err = json.Unmarshal(rr.([]byte), &result)
	if err != nil {
		sendError(c, 9999, err.Error())
	}
	sendOk(c, result)
}
