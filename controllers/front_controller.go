package controllers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// FrontController
type FrontController struct {
	tpl *template.Template
}

// NewFrontController: Constructor of FrontController
func NewFrontController(t *template.Template) *FrontController {
	return &FrontController{t}
}

// Methods have to be capitalized to be exported, eg,
func (frontController FrontController) HomePage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.WriteHeader(http.StatusOK) // 200

	frontController.tpl.ExecuteTemplate(w, "index.html", nil)
}
