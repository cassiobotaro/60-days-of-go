package main

import "net/http"

func healthcheck(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == 200, nil
}
