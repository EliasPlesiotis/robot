package main

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"

)


//func blocks(w http.ResponseWriter, r *http.Request) {
//	t, err := template.ParseFiles("frontend/index.html")
//	log.Print(err.Error())
//	t.Execute(w, "hello")
//}

func main() {
	r := mux.NewRouter()
	
	//	App
	/////////////////////////////////////////////////////////////

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend")))
	
	/////////////////////////////////////////////////////////////

	//http.Handle("/", r)

	go Api()

	log.Fatal(http.ListenAndServe(":8080", r))

}
