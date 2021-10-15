package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"notebook/router"
	"testing"
)

func TestAdminRoute(t *testing.T) {
	gin.SetMode("test")
	ts := httptest.NewServer(router.SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/admin/index.html", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}
