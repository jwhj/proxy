package handler

import (
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	http.HandleFunc("/", H)
	http.ListenAndServe(":8082", nil)
}
