package models

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	key64 = []byte("WXXgBI2Zj5/RCrk8CLmuNPKJ09VP9deXfG3CGoimz5dM9KaNZUnWofK1IAdEdiqbNGApKsW2m/xxkZR5yWXLWw==")
	key32 = []byte("Bpah7x5S6/GdwQ1UP2c5tg+2Pgt9vNovZ/bFjjIxLUE=")
	cookie = securecookie.New(key64, key32)
)

// User
//////////////////////////////////////////
type User struct {
	Username string
	Password string
	Email string
}

//////////////////////////////////////////

func WriteCred() error {	
	f, err := os.Open("credentials.json")
	if err != nil {
		f, err := os.Create("credentials.json")
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Println("create")
		err = json.NewEncoder(f).Encode([][]byte{key64, key32})
	
		return err
	
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	
	var keys [][]byte

	err = json.Unmarshal(bytes, keys)

	cookie = securecookie.New(keys[0], keys[1])

	return err
}

func CreateSession(u *User, w http.ResponseWriter) {
	value := map[string]string{
		"name": u.Username,
		"password": u.Password}

	if encval, err := cookie.Encode("session", value); err == nil {
		http.SetCookie(w, &http.Cookie{Name: "session", Value: encval, Path:"/"})
	} else {
		fmt.Println(err)
	}
}

func GetUsername(r *http.Request) string {
	var username string
	
	if c, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err := cookie.Decode("session", c.Value, &value); err == nil {
			username = value["name"]  
		}
	}	
	
	return username
}

func ClearSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: "session", Value: "", Path: "/", MaxAge: -1})
}