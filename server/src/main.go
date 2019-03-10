package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
	"encoding/json"

	"github.com/gorilla/mux"

	"models"
)


var (
	c models.Commands
)

func files(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c)
}

func folder(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method == "POST" {
		err = models.CreateFile(r, &c)
	} else if r.Method == "GET" {
		err = models.LoadFile(r, &c)
	} else if r.Method == "DELETE" {
		err = models.DeleteFile(r)
	}

	if err != nil {
		log.Print(err)
	}
}


func command(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.CreateCommand(r)
	} else if r.Method == "DELETE" {
		c.DeleteCommand(r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func system(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("system.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	files, err := models.ReadFiles()
	if err != nil {
		fmt.Println(err.Error())
	}

	t.Execute(w, struct{Files []models.File}{files})
}


func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/command/{Id}+{Dir}+{Duration}+{Speed}", command).Methods("POST", "DELETE")
	r.HandleFunc("/files", files).Methods("GET")
	r.HandleFunc("/folder/{Name}", folder).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/system", system).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
