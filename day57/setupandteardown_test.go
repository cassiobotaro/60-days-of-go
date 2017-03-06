package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

func handlerOK(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "error")
}

var server *http.Server

func mySetupFunction() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", handlerOK)
	mux.HandleFunc("/error", handlerError)
	server = &http.Server{
		Addr:           "127.0.0.1:4000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go server.ListenAndServe()
}

func myTeardownFunction() {
	server.Close()
}

func TestHealthcheck(t *testing.T) {
	urlBase := fmt.Sprintf("http://%s", server.Addr)
	_, err := healthcheck(urlBase + "/ok")
	if err != nil {
		t.Errorf("Healthcheck: errors found when not expected ->  %s", err)
	}
	ok, _ := healthcheck(urlBase + "/error")
	if ok {
		t.Errorf("Healthcheck: errors not found when expected")
	}
	_, err = healthcheck("/err")
	if err == nil {
		t.Errorf("Healthcheck: errors not found when expected")
	}
}

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}
