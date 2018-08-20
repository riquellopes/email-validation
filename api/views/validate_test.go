package views

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Route() http.Handler {
	r := mux.NewRouter()
	return r
}

func Test_200(t *testing.T) {
	request, _ := http.NewRequest("GET", "/validate/contato@henriquelopes.com.br", nil)

	record := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/validate/{email}", EmailHandler).Methods("GET")
	route.ServeHTTP(record, request)

	assert.Equal(t, record.Code, http.StatusOK)
	assert.Equal(t, record.Header().Get("Content-Type"), "application/json")
}
