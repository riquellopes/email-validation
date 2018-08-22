package views

import (
	"net/http"
)

// ResquestData -
type ResquestData struct {
	Emails []string `json:"emails"`
}

// ResponseData -
type ResponseData struct{}

// EmailHandler -
func EmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
