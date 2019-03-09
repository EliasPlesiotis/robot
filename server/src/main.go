package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"net"
	"html/template"
	"strconv"
	"os"

	"github.com/gorilla/mux"

	"users"
)


var (
	c Commands
	current_id = 3
	channel = make(chan bool)
	available = false
	begin = false
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(c)
	}
}


func command(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if r.Method == "POST" {
		if params["Dir"] == "{}" && params["Duration"] == "{0}" {
			channel <- true
		} else if string(params["Duration"][1]) == "-" {
			channel <- false
			os.Exit(0)
		} else {
			current_id++
			dur, _ := strconv.Atoi(params["Duration"][1:len(params["Duration"])-1])
			dir := params["Dir"][1:len(params["Dir"])-1]
			c = append(c, &Command{current_id, string(dir), dur})
		}
		
	} else if r.Method == "DELETE" {
		id, _ := strconv.Atoi(params["Id"][1:len(params["Id"])-1])
		for i, com := range c {
			if id == com.Id {
				c = append(c[:i], c[i+1:]...)
			}
		}
	}
}

func ui(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

///////////////////////////////////////////////////
func insider() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		if <-channel {
			fmt.Fprintf(conn, "start")
		} else {
			fmt.Fprintf(conn, "terminate")
		}
	}

}
///////////////////////////////////////////////////


func main() {
	fmt.Println(users.CreateUser())

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/command/{Id}+{Dir}+{Duration}", command).Methods("POST", "DELETE")
	r.HandleFunc("/ui", ui).Methods("GET")
	http.Handle("/", r)

	go insider()

	log.Fatal(http.ListenAndServe(":8080", r))
}