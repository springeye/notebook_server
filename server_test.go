package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"notebook/resources"
	"testing"
)

func TestPageRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()
	t.Run("Admin", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/admin/index.html", ts.URL))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		assert.Equal(t, 200, resp.StatusCode)
	})
}
func TestRoute(t *testing.T) {
	httptest.NewRecorder()
	ts := httptest.NewServer(setupServer())
	defer ts.Close()
	t.Run("Login", func(t *testing.T) {
		params := resources.UserInput{
			Username: "henjue",
			Password: "E10ADC3949BA59ABBE56E057F20F883E",
		}
		paramsByte, _ := json.Marshal(params)
		resp, err := http.Post(fmt.Sprintf("%s/api/user/login", ts.URL), "application/json", bytes.NewBuffer(paramsByte))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		assert.Equal(t, 200, resp.StatusCode)
	})
	t.Run("Get Notebook List", func(t *testing.T) {
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
	})
}
