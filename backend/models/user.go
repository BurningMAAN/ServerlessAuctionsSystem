//go:build !test
// +build !test

package models

type User struct {
	ID       string
	UserName string
	Password string
	Email    string
	Role     string
	Credit   float64
}

type UserUpdate struct {
	UserName string
	Password *string
	Email    *string
	Credit   *float64
}
