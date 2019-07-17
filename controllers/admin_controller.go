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

type AdminController struct{}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (adminController AdminController) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	tpl.ExecuteTemplate(w, "admin/index.html", nil)
}
