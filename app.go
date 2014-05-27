/* Basics of a Webapp using Gorilla & HTML Templates */
package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/index/{name}", index).Methods("GET")
	http.Handle("/", rtr)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := map[string]string{"name": params["name"]}
	t, err := template.ParseFiles("templates/index.html")
	log.Println(err)
	t.Execute(w, name)
}
