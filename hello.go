package main

import (
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// Main function
func main() {

	start := time.Now()
	content, err := ioutil.ReadFile("angularjs.js")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	res1 := strings.ReplaceAll(text, "var", "let")

	// write the whole body at once
	err = ioutil.WriteFile("result_go.js", []byte(res1), 0644)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	log.Printf("%s", elapsed)
	// 6ms
}
