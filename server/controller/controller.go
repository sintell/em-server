package controller

import (
	"net/http"
)

type BaseController struct {
	resource string
	handler  func(http.ResponseWriter, *http.Request)
}

type Controller interface {
	Resourse() string
	Handler() func(http.ResponseWriter, *http.Request)
	Bind() (string, func(http.ResponseWriter, *http.Request))
}

func (b *BaseController) Resourse() string {
	return b.resource
}

func (b *BaseController) Handler() func(http.ResponseWriter, *http.Request) {
	return b.handler
}

func (b *BaseController) Bind() (string, func(http.ResponseWriter, *http.Request)) {
	return b.resource, b.handler
}

func Default(resourse interface{}, handler interface{}) *BaseController {
	r, rOk := resourse.(string)
	h, hOk := handler.(func(http.ResponseWriter, *http.Request))

	if rOk && hOk {
		return &BaseController{r, h}
	} else {
		return &BaseController{"/", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not Found", http.StatusNotFound)
		}}
	}
}
