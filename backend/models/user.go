//go:build !test
// +build !test

package models

type User struct {
	ID       string
	UserName string
	Password string
	Email    string
}
