package main

import (
	"avtoru/app"
	"avtoru/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/api/announcement", controllers.GetAnnById).Methods("GET")
	router.HandleFunc("/api/announcements", controllers.GetAnns).Methods("GET")
	router.HandleFunc("/api/announcements/add", controllers.AddAn).Methods("POST")
	router.HandleFunc("/api/announcement", controllers.DelAn).Methods("DELETE")
	router.HandleFunc("/api/announcement/put", controllers.DelAn).Methods("PUT")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
