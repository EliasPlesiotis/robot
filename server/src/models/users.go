package models

import (

)

// User
//////////////////////////////////////////
type User struct {
	Username string
	Password string
	Email string
}

func CreateUser() error {
	return nil
}
//////////////////////////////////////////