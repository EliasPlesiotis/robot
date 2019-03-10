package models

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	params map[string]string
	current_id = 0
)

//	Command
/////////////////////////////////////
type Command struct {
	Id 		 int    `json:"Id"`
	Dir      string `json:"Dir"`
	Duration int 	`json:"Duration"`
	Speed    int 	`json:"Speed"`
}

/////////////////////////////////////


//	Slice of commands
/////////////////////////////////////
type Commands []Command

func (c *Commands) CreateCommand(r *http.Request) {
	params = mux.Vars(r)
	current_id++
	dur, _ := strconv.Atoi(params["Duration"][1:len(params["Duration"])-1])
	dir := params["Dir"][1:len(params["Dir"])-1]
	speed, _ := strconv.Atoi(params["Speed"][1:len(params["Speed"])-1])
	com := Command{Id: current_id, Dir: dir, Duration: dur, Speed: speed}
	(*c) = append((*c), com)
}

func (c *Commands) DeleteCommand(r *http.Request) {
	params = mux.Vars(r)
	id, _ := strconv.Atoi(params["Id"][1:len(params["Id"])-1])
		for i, com := range *c {
			if id == com.Id {
				(*c) = append((*c)[:i], (*c)[i+1:]...)
			}
		}
}

/////////////////////////////////////
