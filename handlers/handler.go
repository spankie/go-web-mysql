package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/spankie/go-web-mysql/db"
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
	user := User{}
	r.ParseForm()
	user.Email = r.Form.Get("email")
	user.Password = r.Form.Get("password")
	log.Println(user)
	if len(user.Email) < 1 || len(user.Password) < 1 {
		log.Println(err)
		message := "Username or password incorrect"
		err = t.Execute(w, message) // merge.
		if err != nil {
			log.Println(err)
		}
		// status := http.StatusUnauthorized
		return
	}

	// perform a db.Query insert
	insert, err := db.DB.Query(fmt.Sprintf("INSERT INTO users VALUES ( NULL, '%s', '%s' )", user.Email, user.Password))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	err = t.Execute(w, user) // merge.
	if err != nil {
		log.Println(err)
	}
}
