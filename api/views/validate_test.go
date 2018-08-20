package views

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_200(t *testing.T) {
	var req = `{"emails":["contato@henriquelopes.com.br"]}`
	var data = []byte(req)

	request, _ := http.NewRequest("POST", "/validate/", bytes.NewBuffer(data))

	record := httptest.NewRecorder()
	handler := http.HandlerFunc(EmailHandler)
	handler.ServeHTTP(record, request)

	assert.Equal(t, record.Code, http.StatusOK)
	assert.Equal(t, record.Header().Get("Content-Type"), "application/json")
}
