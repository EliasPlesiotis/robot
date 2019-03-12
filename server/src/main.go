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
	u []models.User
)

func files(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c)
}

func folder(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method == "POST" {
		err = models.CreateFile(r, &c)
	} else if r.Method == "GET" {
		params := mux.Vars(r)
		if params["Name"] != "{all}" {	
			err = models.LoadFile(r, &c)
		} else {
			f, err := models.ReadFiles()
			if err != nil {
				log.Print(err)
			}
			json.NewEncoder(w).Encode(f)			
		}
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
	_, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	t, err := template.ParseFiles("../tmpl/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func system(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	username := models.GetUsername(r)
	models.UseLocation(username)

	t, err := template.ParseFiles("../tmpl/system.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("../tmpl/login.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("uname")
		password := r.FormValue("password")

		for _, user := range u {
			if username == user.Username && password == user.Password {
				models.CreateSession(&user, w)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("../tmpl/register.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("uname")
		email := r.FormValue("email")
		password := r.FormValue("password")
		user := models.User{Username: username, Email: email, Password: password}
		u = append(u, user)

		models.CreateSession(&user, w)
		models.GenerateLocation(username)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	models.ClearSession(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/command/{Id}+{Dir}+{Duration}+{Speed}", command).Methods("POST", "DELETE")
	r.HandleFunc("/files", files).Methods("GET")
	r.HandleFunc("/folder/{Name}", folder).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/system", system).Methods("GET")
	r.HandleFunc("/login", login).Methods("GET", "POST")
	r.HandleFunc("/register", register).Methods("GET", "POST")
	r.HandleFunc("/logout", logout)

	http.Handle("/", r)
	
	//fmt.Println(models.WriteCred())

	log.Fatal(http.ListenAndServe(":8080", r))
}
