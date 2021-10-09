package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed assets
var assets embed.FS

func main() {

	//go:embed "assets/BEYOND.pdf"
	var b []byte
	fmt.Println(len(b))

	fmt.Println("http://localhost:8080/assets/index.html")
	fs := http.FileServer(http.FS(assets))
	http.ListenAndServe(":8080", fs)
}
