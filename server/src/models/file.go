package models

import (
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
)


var location string = "EliasPlesiotis"

type File struct {
	Name string `json:Name`
	Url  string `json:Url`
}

func GenerateLocation(username string) error {
	var err error

	os.Chdir("../")
	defer os.Chdir("src")

	err = os.Mkdir(username, 666)
	location = username

	return err
}

func UseLocation(username string) {
	location = username
}

func CreateFile(r *http.Request, c *Commands) error {
	var err error
	
	os.Chdir("../"+ location)
	defer os.Chdir("../src")

	params := mux.Vars(r)
	if err := os.Remove(params["Name"] + ".json"); err != nil {}

	f, err := os.Create(params["Name"] + ".json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(c)

	return err
}

func LoadFile(r *http.Request, c *Commands) error {
	var err error 

	os.Chdir("../"+ location)
	defer os.Chdir("../src")

	params := mux.Vars(r)
	f, err := os.Open(params["Name"][1:len(params["Name"])-1])
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(bytes, c)

	return err
}


func DeleteFile(r *http.Request) error {
	var err error

	os.Chdir("../"+ location)
	defer os.Chdir("../src")

	params := mux.Vars(r)
	err = os.Remove(params["Name"][1:len(params["Name"])-1])

	return err
}


func ReadFiles() ([]File, error) {
	var files []File
	f, err := os.Open("../"+ location)
	if err != nil {
		return files, err
	}

	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}
   
   	for _, file := range fileInfo {
		files = append(files, File{file.Name(), "/folder/{" + file.Name() + "}"})
	}

	return files, nil
}
