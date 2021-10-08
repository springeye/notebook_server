package resources

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Opt      string `json:"opt,omitempty"`
}
type UserResource struct {
	Db *gorm.DB
}

func (r UserResource) Login(context *gin.Context) {
	var user UserInput
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
	}
}
