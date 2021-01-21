package models

import (
	"errors"
	"fmt"
)

// User - struct to parse user body to validation
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

//map to store users
var userMap = make(map[string]User)

// get user details from username
func GetUser(userName string) (*User, error) {
	user, ok := userMap[userName]
	if !ok {
		return nil, errors.New("Failed to fetch user details for username:" + userName)
	}
	return &user, nil
}

//insert new user in map userMap and also check if the user alsready exists
func SetUser(userName, password string) error {
	_, ok := userMap[userName]
	if !ok {
		userMap[userName] = User{UserName: userName, Password: password}
		return nil
	}
	return fmt.Errorf("Cant add Username : %v, as it already exists", userName)
}

//return empty user to Unmarshal
func NewUser() User {
	return User{}
}
