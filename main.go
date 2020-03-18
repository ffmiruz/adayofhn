package main

import (
	"bytes"
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
	b := bytes.Replace(dataBytes, []byte(`href="`), []byte(`href="https://news.ycombinator.com/`), -1)

	dt := time.Now().Hour()
	filename := "site/" + strconv.Itoa(dt) + ".html"
	ioutil.WriteFile(filename, b, 0644)
}
