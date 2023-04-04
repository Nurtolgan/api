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
// @Description Insert a new user
// @ID CreateUserHandler
// @Tags Insert a new user
// @Param request body mongo.Cv true "body json"
// @Accept  json
// @Produce  plain
// @Success 200 {string} string "User created"
// @Failure 400,404 {string} string "error"
// @Router /insertUser [post]
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
// @Description Show a user by username
// @ID showAUser
// @Tags Show a user
// @Param request body mongo.Cv true "username"
// @Produce plain
// @Success 200 {string} string "success"
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
// @Description Delete a user by id
// @ID deleteUser
// @Tags Delete user
// @Param request body mongo.Cv true "json"
// @Produce plain
// @Success 200 {string} string "success"
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
// @Description Update a user by id
// @ID Update user
// @Tags Update user
// @Param request body mongo.Cv true "json"
// @Produce plain
// @Success 200 {string} string "success"
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

// @Summary Show Users with filter
// @Description Get all users or by query params (username, city, birthdaydate, careerobjective)
// @ID GetAllCvsByQuery
// @Tags Show users
// @Param username query string false "username"
// @Success 200 {string} []mongo.Cv "success"
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
