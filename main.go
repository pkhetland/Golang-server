package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

// Welcome struct to hold content for web page
type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Petter", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("template/welcome-template.html"))

	http.Handle("/myPath/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
