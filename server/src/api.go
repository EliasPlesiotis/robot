package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"models/files"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//	Helper Functions
/////////////////////////////////////////////////////////////

//	Error Handling
func errorHandler(view http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if ok := recover(); ok != nil {
				log.Print(ok)
				w.WriteHeader(500)
			}
		}()
		log.Print(r.Method)
		view(w, r)
	}
}

func deleteFiles() error {

	return nil
}

/////////////////////////////////////////////////////////////

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "api")
}

func allfiles(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		all_files, _ := files.GetFiles()
		json.NewEncoder(w).Encode(all_files)
	} else if r.Method == "DELETE" {
		deleteFiles()
	}
}

func file(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(files.GetFile(r))
	} else if r.Method == "POST" {
		files.AddFile(r)
	} else if r.Method == "DELETE" {
		files.DeleteFile(r)
	}
}

func Api() {

	api := mux.NewRouter()

	//	API
	/////////////////////////////////////////////////////////////
	//	Get all files
	api.HandleFunc("/files", allfiles).Methods("GET", "DELETE")
	//	Get a specific file
	api.HandleFunc("/file/{name}", file).Methods("GET", "POST", "DELETE")
	/////////////////////////////////////////////////////////////

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(api)))

}
