package files

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"models/commands"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var ()

//	File Struct
/////////////////////////////////////
type File struct {
	Name     string             //`json:Name`
	Commands []commands.Command //`json:Commands`
}

//	Functions
/////////////////////////////////////

func AddFile(r *http.Request) {
	database, _ := sql.Open("sqlite3", "../db.sqlite3")
	defer database.Close()

	params := mux.Vars(r)
	var f File

	json.NewDecoder(r.Body).Decode(&f)
	fmt.Println(f)

	_, err := database.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", strings.FieldsFunc(params["name"], func(r rune) bool { return r == []rune("{")[0] || r == []rune("}")[0] })[0]))
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = database.Exec("CREATE TABLE IF NOT EXISTS " + strings.FieldsFunc(params["name"], func(r rune) bool { return r == []rune("{")[0] || r == []rune("}")[0] })[0] + " ( Diraction TEXT, Duration INTEGER, Speed INTEGER)")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, com := range f.Commands {
		_, err = database.Exec(fmt.Sprintf("INSERT INTO " + strings.FieldsFunc(params["name"], func(r rune) bool { return r == []rune("{")[0] || r == []rune("}")[0] })[0] + " ( Diraction, Duration, Speed ) VALUES ( '" + com.Diraction + "' , " + com.Duration.(string) + " , " + com.Speed.(string) + " )"))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func DeleteFile(r *http.Request) {
	database, _ := sql.Open("sqlite3", "../db.sqlite3")
	defer database.Close()

	params := mux.Vars(r)

	statement, _ := database.Prepare("DROP TABLE " + strings.FieldsFunc(params["name"], func(r rune) bool { return r == []rune("{")[0] || r == []rune("}")[0] })[0])
	statement.Exec()
}

func GetFile(r *http.Request) []commands.Command {
	database, _ := sql.Open("sqlite3", "../db.sqlite3")
	defer database.Close()

	var com []commands.Command
	params := mux.Vars(r)

	rows, _ := database.Query("SELECT Diraction, Duration, Speed FROM " + strings.FieldsFunc(params["name"], func(r rune) bool { return r == []rune("{")[0] || r == []rune("}")[0] })[0])

	var dir string
	var dur int
	var speed int

	for rows.Next() {
		rows.Scan(&dir, &dur, &speed)
		com = append(com, commands.Command{Diraction: dir, Duration: dur, Speed: speed})
	}

	return com
}

func GetFiles() ([]string, error) {
	database, _ := sql.Open("sqlite3", "../db.sqlite3")
	defer database.Close()

	var tables []string
	rows, err := database.Query("SELECT name FROM sqlite_master")

	for rows.Next() {
		var t string
		rows.Scan(&t)
		tables = append(tables, t)
	}

	return tables, err
}

/////////////////////////////////////
