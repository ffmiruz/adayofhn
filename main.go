package main

import (
	"io/ioutil"
	"time"
)

func main() {
	dt := time.Now()
	// Get byte data to write to file.
	dataString := "Hello friend"
	dataBytes := []byte(dataString)

	// Use WriteFile to create a file with byte data.
	ioutil.WriteFile(dt.Format("01-02-2006 15:04:05 Mon"), dataBytes, 0644)
}
