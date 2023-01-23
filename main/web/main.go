package main

import (
	"go_module/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func main(){
	mux:=chi.NewRouter()
	//mux.Use(middleware.Logger)
	mux.Use(WriteToConsole)

	mux.Get("/",routes.Home)
	mux.Get("/about",routes.About)

	http.ListenAndServe(":8080",mux)

}