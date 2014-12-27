package controller

import (
	"net/http"
)

type BaseController struct {
	resource string
	handler  func(http.ResponseWriter, *http.Request)
}

type Controller interface {
	New() error
}

func New() error {
	return nil
}
