package handlers

import (
	"html/template"
	"net/http"
)

const (
	indexPage      = "index.html"
	productPage    = "product.html"
	addProductPage = "addProduct.html"
)

// Handler Handler
type Handler interface {
	Index(w http.ResponseWriter, r *http.Request)
	//ViewProducts(w http.ResponseWriter, r *http.Request)
	ViewProduct(w http.ResponseWriter, r *http.Request)
	AddProduct(w http.ResponseWriter, r *http.Request)
	//UpdateViewProducts(w http.ResponseWriter, r *http.Request)
}

// SiteHandler SiteHandler
type SiteHandler struct {
	Templates *template.Template
}

// New New
func (h *SiteHandler) New() Handler {
	return h
}
