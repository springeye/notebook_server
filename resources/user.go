package resources

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eko/gocache/v2/store"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"
	. "notebook/cache"
	"notebook/database"
	"notebook/model"
	"time"
)

type UserLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"len=32"`
	Opt      string `json:"opt,omitempty"`
}
type UserRegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"len=32"`
}
type UserResource struct {
	Db    *gorm.DB
	Redis *redis.Client
}

func (r UserResource) Register(context *gin.Context) {
	logger := log.WithFields(log.Fields{})
	var input UserRegisterInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		logger.Debug(err)
		sendFieldError(context, err.Error())
		return
	} else {
		var user model.User
		err = r.Db.Where("username = ?", []string{input.Username}).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//1.input password is encrypted through MD5
			//2.Encrypt "password+salt" via m5d
			//3.Compare passwords stored in the database
			salt := uuid.NewString()
			text := fmt.Sprintf("%s%s", input.Password, salt)
			h := md5.New()
			io.WriteString(h, text)
			md5Pwd := hex.EncodeToString(h.Sum(nil))
			user = model.User{
				Username: input.Username,
				Password: md5Pwd,
				Salt:     salt,
			}
			result := r.Db.Create(&user)
			if result.Error != nil {
				context.AbortWithError(http.StatusInternalServerError, result.Error)
			} else if result.RowsAffected > 0 && user.ID > 0 {
				token := uuid.NewString()
				j, _ := json.Marshal(&user)

				err := Cache.Set(database.RedisContext, fmt.Sprintf("token:%s", token), j, &store.Options{Expiration: time.Duration(-1)})
				if err != nil {
					logger.Panic(err)
					context.AbortWithStatus(http.StatusInternalServerError)
					return
				}
				sendOk(context, gin.H{
					"token": token,
				})
			} else {
				sendError(context, ERROR_REG_ERROR, "register failed")
				return
			}

		} else {
			sendError(context, ERROR_ALREADY_EXIST, "username already exists")
			return
		}
	}
}
func (r UserResource) Login(context *gin.Context) {
	logger := log.WithFields(log.Fields{})
	var input UserLoginInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		logger.Debug(err)
		sendFieldError(context, err.Error())
		return
	} else {
		var result model.User
		err = r.Db.Where("username = ?", []string{input.Username}).First(&result).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(context, ERROR_USER_NOT_FOUND, "user not found")
			return
		} else if err != nil {
			logger.Panic(err)
			sendError(context, ERROR_LOGIN_ERROR, err.Error())
			return
		}
		//1.input password is encrypted through MD5
		//2.Encrypt "password+salt" via m5d
		//3.Compare passwords stored in the database
		text := fmt.Sprintf("%s%s", input.Password, result.Salt)
		h := md5.New()
		io.WriteString(h, text)
		md5Pwd := hex.EncodeToString(h.Sum(nil))
		if md5Pwd != result.Password {
			logger.Debug("username or password is error")
			sendError(context, ERROR_WRONG_USER_OR_PWD, "username or password is error")
			return
		}
		token := uuid.NewString()
		j, _ := json.Marshal(&result)

		err := Cache.Set(database.RedisContext, fmt.Sprintf("token:%s", token), j, &store.Options{Expiration: time.Duration(-1)})
		if err != nil {
			logger.Panic(err)
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		sendOk(context, gin.H{
			"token": token,
		})
	}
}
