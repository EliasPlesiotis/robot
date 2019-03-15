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
	src string
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
			if err != nil {
				fmt.Fprint(w, http.StatusNotFound)
			}
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
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	username, _ := models.GetNamePassword(c.Value)
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

		err := models.LoadUsers(&u)
		if err != nil {
			log.Print(err)
		}

		for _, user := range u {
			if username == user.Username && password == user.Password {
				models.CreateSession(&user, w)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
	} else if r.Method == "DELETE" {
		for i, user := range u {
			c, _ := r.Cookie("session")
			username, password := models.GetNamePassword(c.Value)
			if username == user.Username && password == user.Password {
				u = append(u[:i], u[i+1:]...)
				models.ClearSession(w)
				models.DeleteLocation(username)
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

		for _, user := range u {
			if username == user.Username {
				http.Error(w, "User already exists", http.StatusInternalServerError)
				return
			}
		}

		user := models.User{Username: username, Email: email, Password: password}
		
		err := user.SaveUser()
		if err != nil {
			log.Print(err)
		}

		models.CreateSession(&user, w)
		models.GenerateLocation(username)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	models.ClearSession(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func controller(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	t, err := template.ParseFiles("../tmpl/controller.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func code(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("../tmpl/code.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		src = r.FormValue("code")
		fmt.Println(src)
		http.Redirect(w, r, "/code", http.StatusSeeOther)
	}
}

func view(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(src)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/command/{Id}+{Dir}+{Duration}+{Speed}", command).Methods("POST", "DELETE")
	r.HandleFunc("/files", files).Methods("GET")
	r.HandleFunc("/folder/{Name}", folder).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/system", system).Methods("GET")
	r.HandleFunc("/login", login).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/register", register).Methods("GET", "POST")
	r.HandleFunc("/controller", controller).Methods("GET", "POST")
	r.HandleFunc("/code", code).Methods("GET", "POST")
	r.HandleFunc("/view", view).Methods("GET")
	r.HandleFunc("/logout", logout)

	http.Handle("/", r)
	
	fmt.Println(models.WriteCred())
	err := models.LoadUsers(&u)
	log.Print(err)
	

	log.Fatal(http.ListenAndServe(":8080", r))
}
