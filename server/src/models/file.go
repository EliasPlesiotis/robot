package models

import (
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type File struct {
	Name string `json:Name`
}

func CreateFile(r *http.Request, c *Commands) error {
	var err error

	params := mux.Vars(r)
	if _, err := os.Open(params["Name"] + ".json"); err == nil {
		err = os.Remove(params["Name"] + ".json")
		return err
	}

	f, err := os.Create(params["Name"] + ".json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(c)

	return err
}

func GetFile(r *http.Request, c *Commands) error {
	var err error 

	params := mux.Vars(r)
	f, err := os.Open(params["Name"] + ".json")
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &c)

	return err

}


func DeleteFile(r *http.Request) error {
	var err error

	params := mux.Vars(r)
	err = os.Remove(params["Name"] + ".json")

	return err
}