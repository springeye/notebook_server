package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNotebookList(t *testing.T) {
	gin.SetMode("test")
	httptest.NewRecorder()
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/notebook/list", ts.URL), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"84c767e5-e819-42cb-acc0-a3e695056c04"},
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	assert.Equal(t, 401, resp.StatusCode)
}
