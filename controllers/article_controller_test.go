package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestArticleListHandler(t *testing.T) {
	// define test cases
	tests := []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "a", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("https://localhost:8080/article/list?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder() // First argument to pass to the handler
			aCon.ArticleListHandler(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("expected: %d, but got: %d", tt.resultCode, res.Code)
			}
		})
	}
}

func TestArticleDetailHandler(t *testing.T) {
	// define test cases
	tests := []struct {
		name       string
		id         string
		resultCode int
	}{
		{name: "number pathparam", id: "1", resultCode: http.StatusOK},
		{name: "alphabet pathparam", id: "a", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("https://localhost:8080/article/%s", tt.id)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder() // First argument to pass to the handler
			r := mux.NewRouter()
			r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req) // send request to the router for testing

			if res.Code != tt.resultCode {
				t.Errorf("expected: %d, but got: %d", tt.resultCode, res.Code)
			}
		})
	}
}
