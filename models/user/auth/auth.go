package auth

import (
	"ExCloud/models/user/register"
)

type UserAuth struct {
	register.UserRegistration        //get login + pass
	Token                     string //get token
}

type Authorization interface {
	SignIn(login, password string, auth UserAuth) error //user not found
	SignUp(login, password string) (UserAuth, error)    //user already exists
	CreateToken(login, password string)
	GetToken()
	ValidateToken()
}

/*type User struct {
	Login    string
	Password string
}

func CreateUser(login, password string) (User, error) {
	stmt := `insert into user value login,password where login=$1 and password=$2`
	func() {

	}()
}

func GetUserFromDb() string {
	var login string
	var password string
	stmt := `select from user value login,password`
	func() {

	}()
	return login, password
}

func (u UserAuth) SignIn(login, password string) {

	if u.Login != "" && u.Password != "" {

	}
	u.Password
}
func SignUp(login, password string) {

}*/
