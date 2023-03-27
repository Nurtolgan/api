package main

import (
	"api/debugger"
	"api/mongo"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func readGetResponse(url string) []byte {
	response, err := http.Get(url)
	debugger.CheckError("Get", err)
	b, err := io.ReadAll(response.Body)
	debugger.CheckError("Read Body", err)
	return b
}

func readPostResponse(url string, contentType string, body io.Reader) []byte {
	response, err := http.Post(url, contentType, body)
	debugger.CheckError("Post", err)
	b, err := io.ReadAll(response.Body)
	debugger.CheckError("Read Body", err)
	return b
}

func readDeleteResponse(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	debugger.CheckError("New Request", err)
	response, err := client.Do(req)
	debugger.CheckError("Do", err)
	b, err := io.ReadAll(response.Body)
	debugger.CheckError("Read Body", err)
	return b
}

func readUpdateResponse(url string, contentType string, body io.Reader) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, body)
	debugger.CheckError("New Request", err)
	req.Header.Set("Content-Type", contentType)
	response, err := client.Do(req)
	debugger.CheckError("Do", err)
	b, err := io.ReadAll(response.Body)
	debugger.CheckError("Read Body", err)
	return b
}

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(index)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "API"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestInsertUser(t *testing.T) {
	cvJSON, err := ioutil.ReadFile("example/example.json")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(insertUser))
	defer ts.Close()
	req, err := http.NewRequest("POST", ts.URL+"/insertUser", bytes.NewBuffer(cvJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestShowAUser(t *testing.T) {
	b := readGetResponse("http://localhost:8000/users/Ray")
	var cv mongo.Cv
	if json.Unmarshal(b, &cv) != nil {
		t.Errorf("Failed to unmarshal json")
	}
}

func TestDeleteUser(t *testing.T) {
	b := readDeleteResponse("http://localhost:8000/delete/641c835be1045e27dafbb105")
	if string(b) != "" {
		t.Errorf("Failed to delete user")
	}
}

func TestUpdateUser(t *testing.T) {
	cvJSON, err := ioutil.ReadFile("example/example.json")
	if err != nil {
		t.Fatal(err)
	}
	b := readUpdateResponse("http://localhost:8000/update/641c83c2e1045e27dafbb107", "application/json", bytes.NewBuffer(cvJSON))
	if string(b) != "" {
		t.Errorf("Failed to update user")
	}
}
