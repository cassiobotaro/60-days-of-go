package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// HTTPBin represents the body returned by GET request in https://httpbin.org/get
type HTTPBin struct {
	Args struct {
	} `json:"args"`
	Headers struct {
		Accept                  string `json:"Accept"`
		AcceptEncoding          string `json:"Accept-Encoding"`
		AcceptLanguage          string `json:"Accept-Language"`
		Cookie                  string `json:"Cookie"`
		Host                    string `json:"Host"`
		Referer                 string `json:"Referer"`
		SaveData                string `json:"Save-Data"`
		UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests"`
		UserAgent               string `json:"User-Agent"`
	} `json:"headers"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {
	// Create a file at tmp directory
	f, err := os.Create("/tmp/dat2")
	if err != nil {
		log.Fatal(err)
	}
	// the action of close file is deferred(run after function finnish)
	defer f.Close()
	// write some content(remember that file is open until the end of this function)
	f.WriteString("some content")
	// performs a http request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	var bodyMap HTTPBin
	// read the body and decode the content
	err = json.NewDecoder(resp.Body).Decode(&bodyMap)
	// don't forget to close the body
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bodyMap.Origin = %+v\n", bodyMap.Origin)
	fmt.Printf("bodyMap.URL = %+v\n", bodyMap.URL)
}
