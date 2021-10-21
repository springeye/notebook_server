package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log2 "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"notebook/database"
	"notebook/model"
)

type NotebookResource struct {
	Db    *gorm.DB
	Redis *redis.Client
}

// @BasePath /api
// @Summary get notebook list
// @Schemes
// @Description get notebook list
// @Tags notebook
// @Accept json
// @Produce json
// @Success 200 {array} model.Notebook
// @Router /notebook/list [get]
// @Security user_token
func (r *NotebookResource) GetNotebookList(c *gin.Context) {

	var results []model.Notebook
	err := database.Database.Find(&results).Error
	if err != nil {
		log2.Panic(err)
	} else {
		sendOk(c, results)
	}

}
