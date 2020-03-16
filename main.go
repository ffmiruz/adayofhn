package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var client http.Client
	resp, err := client.Get("https://news.ycombinator.com/news")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	dt := time.Now().Hour()
	filename := "site/" + strconv.Itoa(dt) + ".html"
	ioutil.WriteFile(filename, dataBytes, 0644)
}
