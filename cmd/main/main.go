package main

import (
	"log"
	"net/http"

	"github.com/Danelarrate/go-bookstore/cmd/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServer("localhost:9010",r))
}