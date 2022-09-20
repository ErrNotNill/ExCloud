package register

type UserRegistration struct {
	Id       int
	Login    string
	Password string
	//Token         string
	Success       bool
	StorageAccess string
}
type Register interface {
	CreateUser(login, password string) (UserRegistration, error) //error if user already exists
}
