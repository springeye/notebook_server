package resources

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"
	"notebook/db"
	"notebook/model"
	"strings"
)

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"len=32"`
	Opt      string `json:"opt,omitempty"`
}
type UserResource struct {
	Db    *gorm.DB
	Redis *redis.Client
}

func (r UserResource) Login(context *gin.Context) {
	logger := log.WithFields(log.Fields{})
	var input UserInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		logger.Debug(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
	} else {
		var result model.User
		err = r.Db.Where("username = ?", []string{input.Username}).First(&result).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.Status(http.StatusNotFound)
			context.Abort()
		} else if err != nil {
			logger.Panic(err)
			context.Status(http.StatusInternalServerError)
			context.Abort()
		}
		//1.input password is encrypted through MD5
		//2.Encrypt "password+salt" via m5d
		//3.Compare passwords stored in the database
		text := fmt.Sprintf("%s%s", input.Password, result.Salt)
		h := md5.New()
		io.WriteString(h, text)
		md5Pwd := string(h.Sum(nil))
		if !strings.EqualFold(md5Pwd, result.Password) {
			logger.Debug("username or password is error")
			context.Status(http.StatusUnauthorized)
			context.Abort()
		}
		token := uuid.NewString()
		j, _ := json.Marshal(result)
		r.Redis.SAdd(db.RedisContext, fmt.Sprintf("token:%s", token), string(j))
		context.JSON(200, gin.H{
			"code": 0,
			"msg":  "",
			"data": gin.H{
				"token": token,
			},
		})
	}
}
