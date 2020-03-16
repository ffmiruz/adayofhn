package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	url := "https://news.ycombinator.com/news"
	var client http.Client
	resp, err := client.Get(url)
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
