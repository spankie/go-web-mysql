package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

// HomeHandler serves index page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html") // Create a template.
	cwd, _ := os.Getwd()
	p := path.Join(cwd, "public", "index.html")
	t, err := t.ParseFiles(p) // Parse template file.
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, nil) // merge.
	if err != nil {
		log.Println(err)
	}
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html") // Create a template.
	cwd, _ := os.Getwd()
	p := path.Join(cwd, "public", "index.html")
	t, err := t.ParseFiles(p) // Parse template file.
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, nil) // merge.
	if err != nil {
		log.Println(err)
	}
}
