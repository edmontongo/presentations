package main

type User struct {
	Username string `json:"username,omitempty" db:"username"`
	Password string `json:"password,omitempty" db:"password"`
}
