package handler

import (
	"fmt"
	"testing"
)

func TestCheckLoginCase(t *testing.T) {
	var login = "ADMInaZ123"
	v := CheckLoginCase(login)
	if !v {
		fmt.Println(len(login))
		t.Fatal("Incorrect login")
	} else {
		t.Fatal()
	}
}
func TestCheckPasswordCase(t *testing.T) {
	var password = "User123User123@"
	v := CheckPasswordCase(password)
	if !v {
		fmt.Println(len(password))
		t.Error("Incorrect password")
	}
}
