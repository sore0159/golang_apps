package main

import (
	"fmt"
	"image/png"
	"net/http"
)

func ServeHTTP() {
	fmt.Println("Starting HTTP server at 127.0.0.1:8080")
	http.HandleFunc("/", servePng)
	http.HandleFunc("/bez.html", bezHTML)
	http.HandleFunc("/bez.png", bezPNG)
	http.HandleFunc("/bez.js", bezJS)
	http.HandleFunc("/bez.api", bezAPI)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error serving: ", err)
	}
}

func servePng(w http.ResponseWriter, r *http.Request) {
	d, err := LoadDots()
	if err != nil {
		http.Error(w, "load dots error", 500)
		return
	}
	img, err := d.CreateImg()
	if err != nil {
		http.Error(w, "create img error", 500)
		return
	}
	png.Encode(w, img)
	return
}
