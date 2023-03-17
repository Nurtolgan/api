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

	debugger.CheckError("Failed to read request body", err)

	var cv mongo.Cv
	err = json.Unmarshal(body, &cv)
	debugger.CheckError("Failed to parse request body", err)

	err = cv.Validate()
	debugger.CheckError("Invalid request body", err)

	err = mongo.CreateUserHandler(cv)
	debugger.CheckError("Failed to create user", err)
    w.WriteHeader(http.StatusCreated)
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
