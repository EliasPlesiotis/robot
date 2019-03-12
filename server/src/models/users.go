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
	key64 = securecookie.GenerateRandomKey(64)
	key32 = securecookie.GenerateRandomKey(32)
	cookie = securecookie.New(key64, key32)
)

// User
//////////////////////////////////////////
type User struct {
	Username string
	Password string
	Email string
}

func (u User) SaveUser() error {
	var users []User
	
	err := LoadUsers(&users)
	if err != nil {
		return err
	}

	if err = os.Remove("users.json"); err != nil {}
	f, err := os.Create("users.json")

	users = append(users, u)

	err = json.NewEncoder(f).Encode(users)
	
	return err
}
//////////////////////////////////////////

func LoadUsers(u *[]User) error {
	f, err := os.Open("users.json")
	if err != nil {
		_, err := os.Create("users.json")
		if err != nil {
			return err
		}
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	err = json.Unmarshal(b, u)
	
	return err
}

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

	err = json.Unmarshal(bytes, &keys)

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

func GetNamePassword(r string) (string, string) {
	var username string
	var password string
	
	value := make(map[string]string)
	if err := cookie.Decode("session", r, &value); err == nil {
		username = value["name"]  
		password = value["password"]
	}
		
	
	return username, password
}

func ClearSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: "session", Value: "", Path: "/", MaxAge: -1})
}