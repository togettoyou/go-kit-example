package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println(get("http://127.0.0.1:8888/name"))
	log.Println(get("http://127.0.0.1:8888/age"))
}

func get(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
