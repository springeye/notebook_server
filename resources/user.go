package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	} else {
		context.JSON(200, gin.H{

			"you_request": gin.H{
				"username": user.Username,
				"password": user.Password,
			},
		})
	}
}
