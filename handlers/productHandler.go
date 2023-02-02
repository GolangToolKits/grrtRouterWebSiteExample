package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/GolangToolKits/grrt"
)

// Product Product
type Product struct {
	ID   int64
	Sku  string
	Name string
}

// ViewProduct ViewProduct
func (h *SiteHandler) ViewProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pidstr := vars["id"]
	sku := vars["sku"]
	pid, _ := strconv.ParseInt(pidstr, 10, 64)

	fmt.Println("id: ", pid)
	fmt.Println("sku: ", sku)

	var p Product
	p.ID = pid
	p.Sku = sku

	h.Templates.ExecuteTemplate(w, productPage, &p)

}

// AddProduct AddProduct
func (h *SiteHandler) AddProduct(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	sku := r.FormValue("sku")

	var p Product
	p.Name = name
	p.Sku = sku

	h.Templates.ExecuteTemplate(w, addProductPage, &p)

}
