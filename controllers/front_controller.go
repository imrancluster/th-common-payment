package controllers

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// FrontController
type FrontController struct{}

// NewFrontController: Constructor of FrontController
func NewFrontController() *FrontController {
	return &FrontController{}
}

// Methods have to be capitalized to be exported, eg,
func (frontController FrontController) HomePage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.WriteHeader(http.StatusOK) // 200

	tpl.ExecuteTemplate(w, "index.html", nil)
}
