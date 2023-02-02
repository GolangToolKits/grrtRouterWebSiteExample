package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	mux "github.com/GolangToolKits/grrt"
	hd "github.com/GolangToolKits/grrtRouterWebSiteExample/handlers"
)

func main() {

	var sh hd.SiteHandler

	sh.Templates = template.Must(template.ParseFiles("./static/index.html",
		"./static/product.html", "./static/addProduct.html"))

	router := mux.NewRouter()

	h := sh.New()

	router.HandleFunc("/", h.Index).Methods("GET")
	router.HandleFunc("/product/{id}/{sku}", h.ViewProduct).Methods("GET")
	router.HandleFunc("/addProduct", h.AddProduct).Methods("POST")

	port := "8080"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Web UI is running on port 8080!")

	http.ListenAndServe(":"+port, (router))
}

//go mod init github.com/GolangToolKits/grrtRouterWebSiteExample
