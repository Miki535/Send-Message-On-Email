package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
)

// templates:
var tpl = template.Must(template.ParseFiles("templates/index.html"))
var tplp = template.Must(template.ParseFiles("templates/about.html"))

func main() {
	http.HandleFunc("/about", about)

	http.HandleFunc("/", HomeFunc)

	//start server on localhost :8080

	http.ListenAndServe(":8080", nil)
}

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// secret code and own email
		secretKey := r.FormValue("secretkey")
		ownemail := r.FormValue("ownemail")
		email := r.FormValue("email")
		message := r.FormValue("message")

		auth := smtp.PlainAuth("", ownemail, secretKey, "smtp.gmail.com")

		to := []string{email}
		msg := message
		// send message on gmail
		err := smtp.SendMail("smtp.gmail.com:587", auth, ownemail, to, []byte(fmt.Sprint(msg)))

		if err != nil {
			http.Error(w, "ERROR! \n StatusBadRequest", http.StatusBadRequest)
			return
		}

	}
	tpl.Execute(w, nil)
}

//func about for our frontend

func about(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}
	tplp.Execute(w, nil)

}
