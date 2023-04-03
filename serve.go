package main

import (
	"api/debugger"
	"api/mongo"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @Summary index
// @ID index
// @Tags index
// @Produce plain
// @Success 200 {string} string "OK"
// @Failure 400,404 {string} string "error"
// @Router / [get]
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API"))
}

// @Summary insertUser
// @ID insertUser
// @Tags insertUser
// @Produce plain
// @Success 200 {string} string "Created"
// @Failure 400,404 {string} string "error"
// @Router /insertUser [get]
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

// @Summary showAUser
// @ID showAUser
// @Tags showAUser
// @Produce plain
// @Success 200 {string} string "OK"
// @Failure 400,404 {string} string "error"
// @Router /users/{username} [get]
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

// @Summary deleteUser
// @ID deleteUser
// @Tags deleteUser
// @Produce plain
// @Success 200 {string} string "OK"
// @Failure 400,404 {string} string "error"
// @Router /delete/{id} [delete]
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := mongo.DeleteUserById(id)
	debugger.CheckError("Failed to delete user", err)
	w.WriteHeader(http.StatusNoContent)

}

// @Summary updateUser
// @ID updateUser
// @Tags updateUser
// @Produce plain
// @Success 200 {string} string "OK"
// @Failure 400,404 {string} string "error"
// @Router /update/{id} [put]
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

// @Summary showUsers
// @ID showUsers
// @Tags showUsers
// @Produce plain
// @Success 200 {string} string "OK"
// @Failure 400,404 {string} string "error"
// @Router /showusers [get]
func showUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	args := []string{
		query.Get("username"),
		query.Get("city"),
		query.Get("birthdaydate"),
		query.Get("careerobjective"),
	}
	cvs, err := mongo.GetAllCvsByQuery(args...)
	debugger.CheckError("Failed to get users", err)

	jsonBytes, err := json.Marshal(cvs)
	debugger.CheckError("Failed to marshal json", err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func HandleRequest() {
	r := mux.NewRouter()
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/insertUser", insertUser).Methods("POST")
	r.HandleFunc("/users/{username}", showAUser).Methods("GET")
	r.HandleFunc("/delete/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/update/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/showusers", showUsers).Methods("GET")
	debugger.CheckError("Listen and serve", http.ListenAndServe(":8010", handlers.CORS(header, methods, origins)(r)))
}
