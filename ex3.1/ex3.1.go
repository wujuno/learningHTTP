package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	// 문자열로 200 OK
	log.Println("Status:", resp.Status)
	// 수치로 200
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}