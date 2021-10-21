package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetNotebookList(t *testing.T) {
	gin.SetMode("test")
	httptest.NewRecorder()
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()
	params := UserLoginInput{
		Username: "test",
		Password: "E10ADC3949BA59ABBE56E057F20F883E",
	}
	paramsByte, _ := json.Marshal(params)
	resp, err := http.Post(fmt.Sprintf("%s/api/user/login", ts.URL), "application/json", bytes.NewBuffer(paramsByte))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	var result gin.H
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	token := result["data"].(map[string]interface{})["token"].(string)
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/notebook/list", ts.URL), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}
	time.Sleep(time.Second * 7)
	resp2, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	assert.Equal(t, 200, resp2.StatusCode)
}
