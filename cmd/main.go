package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ws/internal/browser"
	"ws/internal/database"
)

var schema string

func main() {
	db, err := database.GetDatabase()

	if err != nil || db.Ping() != nil {
		log.Fatalf("Error create DB: %s", err)
	}

	browserController := new(browser.Controller)
	browserController.Init(db)

	r := mux.NewRouter()

	r.HandleFunc("/", browserController.CreateScreenshot).
		Methods("POST")
	r.HandleFunc("/", browserController.GetAll).
		Methods("GET")
	r.HandleFunc("/image/{id}", browserController.GetImage).
		Methods("GET")

	log.Println("listen :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}
