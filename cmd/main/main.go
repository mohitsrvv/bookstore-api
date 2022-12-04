package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mohit/bookstore-api/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localost:9010", r))
}