package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func JsonRequest(url string, structure interface{}) (err error) {
	// initialize a http client
	client := &http.Client{}

	// create a GET request but don't perform yet
	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, url, nil)
	// errors can occurs when create a request
	// for example the url protocol can be wrong
	if err != nil {
		return err
	}
	// Added content type json
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	// perform request and retrieve a response
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	// don't forget to close the body
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(structure)
	if err != nil {
		return err
	}
	return nil

}

func main() {
	// map json response into struct
	s := struct {
		Origin  string            `json:"origin"`
		Url     string            `json:"url"`
		Args    map[string]string `json:"args"`
		Headers map[string]string `json:"headers"`
	}{}
	// Call JsonRequest and verify raised error
	switch err := JsonRequest("http://httpbin.org/get", &s).(type) {
	// no errors, print Origin(your ip)
	case nil:
		fmt.Printf("Origin: %s\n", s.Origin)
	// there's something wrong
	case *url.Error:
		fmt.Println("Error on request")
	// error while parse json
	case *json.SyntaxError:
		fmt.Printf("Error while parse json: %q", err)
	// possible error on structure mapping
	case *json.InvalidUnmarshalError:
		fmt.Println("Problem while map json to structure")
	// there are many errors that can occurs
	default:
		fmt.Printf("Unexpected error %T = %+v\n", err, err)
	}
	// this approach prevent parse error messages.
	// Even if the error messages change,
	// you continue to capture them correctly
}
