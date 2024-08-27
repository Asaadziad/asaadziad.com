package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	fmt.Println("Serving static files on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
