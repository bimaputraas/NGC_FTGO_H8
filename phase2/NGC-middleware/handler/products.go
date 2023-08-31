package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductHandler struct {
}


func (h ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Add Recipe")
}

func (h ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Delete Recipe")
}

func (h ProductHandler) ViewProducts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "View Recipes")
}

func (h ProductHandler) ViewProductById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "View Recipe by id")
}

func (h ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Update Recipe")
}

