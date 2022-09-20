package models

import (
	"errors"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token"`
}

type Users interface {
	CreateUser(login, password string) error
	GetUserById()
	EditUser()   //todo
	DeleteUser() //todo
}

func (u User) CreateUser(login, password string) error {

	if login == "" {
		return errors.New("login must have some symbols")
	}
	if len(login) < 8 {
		return errors.New("login length must be more than 8 symbols")
	}
	byteconv := []byte(login)
	for i := range byteconv {
		if byteconv[i] == ' ' {
			return errors.New("not allow empty fields")
		}
	}
	return nil
}

// ???
/*func CreateAdmin(token string) {
	stmt := `insert into users value login=admin and password=admin`
	token := `882c6bb9a14d326589cfe6acc6929bd6`
}*/
