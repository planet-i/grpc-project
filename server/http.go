package server

import (
	"net/http"
)

func NewHttp() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
