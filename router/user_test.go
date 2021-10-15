package router

import (
	"bytes"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func TestGenerateOTPCode(t *testing.T) {
	secret := base32.StdEncoding.EncodeToString([]byte("12345678901234567890"))
	fmt.Println(secret)
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		panic(err)
	}
	println(passcode)
}
func TestLogin(t *testing.T) {
	gin.SetMode("test")
	httptest.NewRecorder()
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()
	params := UserLoginInput{
		Username: "henjue",
		Password: "E10ADC3949BA59ABBE56E057F20F883E",
	}
	paramsByte, _ := json.Marshal(params)
	resp, err := http.Post(fmt.Sprintf("%s/api/user/login", ts.URL), "application/json", bytes.NewBuffer(paramsByte))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}
