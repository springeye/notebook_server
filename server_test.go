package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"notebook/resources"
	"testing"
)

func checkResponse(t *testing.T, resp *http.Response) {

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
func TestLoginRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()
	params := resources.UserInput{
		Username: "henjue",
		Password: "alalalalalalalalsjsjdhfgwiwiwiwi",
	}
	paramsByte, _ := json.Marshal(params)
	resp, err := http.Post(fmt.Sprintf("%s/api/user/login", ts.URL), "application/json", bytes.NewBuffer(paramsByte))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	checkResponse(t, resp)
}
func TestNoteBookListRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/notebook/list", ts.URL), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"jajdklfajlkdfjalksdfjlk"},
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	checkResponse(t, resp)
}
