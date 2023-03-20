package main

import (
	"api/debugger"
	"api/mongo"
	"encoding/json"
	"io"
	"net/http"

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

func showAUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	cv, err := mongo.GetCVByUsername(username)
	debugger.CheckError("Failed to get user", err)

	jsonBytes, err := json.Marshal(cv)
	debugger.CheckError("Failed to marshal json", err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := mongo.DeleteUserById(id)
	debugger.CheckError("Failed to delete user", err)
	w.WriteHeader(http.StatusNoContent)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := io.ReadAll(r.Body)
	debugger.CheckError("Failed to read request body", err)

	var cv mongo.Cv
	err = json.Unmarshal(body, &cv)
	debugger.CheckError("Failed to parse request body", err)
	err = cv.Validate()
	debugger.CheckError("Invalid request body", err)

	err = mongo.UpdateUserById(id, cv)
	debugger.CheckError("Failed to update user", err)

	w.WriteHeader(http.StatusNoContent)
}

func HandleRequest() {
	r := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/insertUser", insertUser).Methods("POST")
	r.HandleFunc("/users/{username}", showAUser).Methods("GET")
	r.HandleFunc("/delete/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/update/{id}", updateUser).Methods("PUT")
	http.ListenAndServe(":8080", handlers.CORS(header, methods, origins)(r))

}
