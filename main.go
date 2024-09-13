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
	// Remove the initial repo fetching

	index := views.Index(nil) // Pass nil instead of projects

	http.Handle("/", templ.Handler(index))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Add a new handler for fetching repos
	http.HandleFunc("/fetch-repos", fetchReposHandler)

	fmt.Println("Serving on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func fetchReposHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.github.com/users/Asaadziad/repos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var repos []types.Repo
	if err := json.Unmarshal(body, &repos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projects := components.Projects(repos)
	projects.Render(r.Context(), w)
}
