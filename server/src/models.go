package main

//	Command
/////////////////////////////////////
type Command struct {
	Id 		 int    `json:"Id"`
	Dir      string `json:"Dir"`
	Duration int 	`json:"Duration"`
}
/////////////////////////////////////

//	Slice of commands
/////////////////////////////////////
type Commands []*Command
/////////////////////////////////////