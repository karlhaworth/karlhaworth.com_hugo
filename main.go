// Package main ...
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-rod/rod"
)

func main() {
	// get local file
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	filepath := path + "/public/index.html"

	rod.New().MustConnect().MustPage("file://" + filepath).MustWaitLoad().MustPDF("./public/karl_haworth-resume.pdf")
	fmt.Println("wrote ./public/karl_haworth-resume.pdf")
}
