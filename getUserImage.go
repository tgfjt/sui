package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetUserImage(p Profile) {
	res, err := http.Get(p.Image)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var ext string

	switch http.DetectContentType(body) {
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	default:
		ext = ".jpg"
	}

	filename := "images/" + p.DisplayName + ext

	file, err := os.Create(filename)

	fmt.Printf("download: %v\n", filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(body)
}
