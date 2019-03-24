package commands

import (
	_ "github.com/mattn/go-sqlite3"
)

var (
//database, _ = sql.Open("sqlite3", "../db.sqlite3")
)

//	Command Struct
/////////////////////////////////////
type Command struct {
	Diraction string      //`json:Diraction`
	Duration  interface{} //`json:Duration`
	Speed     interface{} //`json:Speed`
}

//	Functions
/////////////////////////////////////

/////////////////////////////////////
