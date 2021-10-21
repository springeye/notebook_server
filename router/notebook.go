package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	log2 "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"notebook/database"
	"notebook/model"
	"notebook/store"
)

type NotebookUpdateInput struct {
	NotebookCreateInput
	Uuid string `json:"uuid,omitempty"`
}
type NotebookCreateInput struct {
	Title    string `json:"title"`
	Pid      uint   `json:"pid,omitempty"`
	Password string `json:"password,omitempty"`
}
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

// @BasePath /api
// @Summary create a notebook
// @Schemes
// @Description create a notebook
// @Tags notebook
// @Accept json
// @Produce json
// @Success 200 {array} NotebookCreateInput
// @Router /notebook [post]
// @Security user_token
func (r *NotebookResource) Create(c *gin.Context) {
	var input NotebookCreateInput
	c.ShouldBindJSON(&input)
	user := getUser(c)
	notebook := model.Notebook{
		Uuid:     uuid.New().String(),
		UserId:   user.ID,
		Title:    input.Title,
		Pid:      &input.Pid,
		Password: input.Password,
	}
	err := r.Db.Create(&notebook).Error
	if err != nil {
		sendError(c, 500, err.Error())
	} else {
		sendOk(c, nil)
	}

}
func getUser(c *gin.Context) *model.User {
	s := store.Default(c)
	token := c.GetHeader("Authorization")
	val := s.Get(fmt.Sprintf("token:%s", token))
	var user model.User
	err := json.Unmarshal([]byte(val), &user)
	if err == nil {
		return &user
	} else {
		return nil
	}

}
