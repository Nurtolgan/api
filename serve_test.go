package main

import (
	"api/debugger"
	"bytes"
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
