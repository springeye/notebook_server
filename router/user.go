package router

import (
	"crypto/md5"
	"encoding/hex"
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
	"notebook/model"
	"notebook/store"
)

type UserLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,len=32"`
	Opt      string `json:"opt,omitempty"`
}
type AuthOutput struct {
	Token string `json:"token"`
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
			} else if result.RowsAffected > 0 && user.ID != "" {
				token := uuid.NewString()
				j, _ := json.Marshal(&user)
				cache := store.Default(context)
				cache.Set(fmt.Sprintf("token:%s", token), string(j))

				sendOk(context, &AuthOutput{token})
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

// @BasePath /api
// @Summary login
// @Schemes
// @Description user login
// @Tags user
// @Accept json
// @Produce json
// @Param account body UserLoginInput true "login user info"
// @Success 200 {object} AuthOutput
// @Router /user/login [post]
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

		c := store.Default(context)
		c.Set(fmt.Sprintf("token:%s", token), string(j))

		sendOk(context, &AuthOutput{token})
	}
}
