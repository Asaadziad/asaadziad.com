package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Asaadziad/devlog/types"
	"github.com/Asaadziad/devlog/views"
	"github.com/Asaadziad/devlog/views/components"
	"github.com/a-h/templ"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/Asaadziad/repos")
	if err != nil {
		fmt.Println(err)
		return
	}
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var r []types.Repo
	json.Unmarshal(body, &r)

	projects := components.Projects(r)
	index := views.Index(projects)

	http.Handle("/", templ.Handler(index))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Serving static files on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
