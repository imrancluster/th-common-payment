package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// FrontController
type FrontController struct{}

// NewFrontController: Constructor of FrontController
func NewFrontController() *FrontController {
	return &FrontController{}
}

// Methods have to be capitalized to be exported, eg,
func (frontController FrontController) HomePage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", "Welcome to Home page")
}
