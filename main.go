package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	fmt.Println("Serving static files on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
