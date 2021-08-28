package main

import (
	"avtoru/app"
	"avtoru/controllers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	app.AddRoute("/api/user/new", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodPost: controllers.CreateAccount,
	})
	app.AddRoute("/api/user/login", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodPost: controllers.Authenticate,
	})
	app.AddRoute("/api/user", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodGet: controllers.GetAccount,
	})
	app.AddRoute("/api/announcement", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodGet:    controllers.GetAnnById,
		http.MethodDelete: controllers.DelAn,
		http.MethodPut:    controllers.UpdAn,
	})
	app.AddRoute("/api/announcements", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodGet: controllers.GetAnns,
	})
	app.AddRoute("/api/announcements/add", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodPost: controllers.AddAn,
	})

	//router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	//router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	//router.HandleFunc("/api/user", controllers.GetAccount).Methods("GET")
	//router.HandleFunc("/api/announcement", controllers.GetAnnById).Methods("GET")
	//router.HandleFunc("/api/announcements", controllers.GetAnns).Methods("GET")
	//router.HandleFunc("/api/announcements/add", controllers.AddAn).Methods("POST")
	//router.HandleFunc("/api/announcement", controllers.DelAn).Methods("DELETE")
	//router.HandleFunc("/api/announcement/put", controllers.DelAn).Methods("PUT")

	//router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	//err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	//if err != nil {
	//	fmt.Print(err)
	//}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
