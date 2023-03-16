package main

import (
	"net/http"
	"io"
	"api/debugger"
	"api/mongo"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API"))
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	debugger.CheckError("ReadAll", err)

	var cv mongo.Cv
	debugger.CheckError("Unmarshal", json.Unmarshal(body, &cv))
	mongo.CreateUserHandler(cv)
}

func HandleRequest() {
	r := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/insertUser", insertUser).Methods("POST")

	http.ListenAndServe(":8080", handlers.CORS(header, methods, origins)(r))

}
