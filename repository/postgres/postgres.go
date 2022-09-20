package postgres

import (
	"ExCloud/cache"
	"ExCloud/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	db    *sql.DB
	cache cache.Cache
	log   log.Logger
	UserPostgres
}

type UserPostgres interface {
	CreateUser(login, password string) error
	Authenticate(login, password string) error
	GetUserById()
}

func (u *Repository) AuthenticateStorage(login, password string) error {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}
	/*login = u.Login
	password = u.Password*/

	user := models.User{}
	query := `SELECT login,password FROM users WHERE login=$1 AND password=$2`
	err = db.QueryRow(query, &login, &password).Scan(&user.Login, &user.Password)
	if err != nil {
		return err
	}
	//queryToken = ``

	return err
}

//UserExistsOld

func (u *Repository) LoginExists(login string) bool {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}

	var exists bool
	//query := `SELECT EXISTS (SELECT id FROM users WHERE id = $1)`
	qry := `SELECT EXISTS (SELECT login FROM users WHERE login =$1);`
	db.QueryRow(qry, login).Scan(&exists)

	return exists
}
func (u *Repository) UserExists(login, password string) bool {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}

	var exists bool
	//query := `SELECT EXISTS (SELECT id FROM users WHERE id = $1)`
	qry := `SELECT EXISTS (SELECT login,password FROM users WHERE login =$1 AND password=$2);`
	db.QueryRow(qry, &login, &password).Scan(&exists)

	return exists
}

/*func (u *Repository) UserExists(login, password string) {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}
	query := `select login, password from users where login=$1 and password=$2` //returning login
	db.QueryRow(query, login, password)

}*/

func (u *Repository) CreateUser(login, password string) (int, error) {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}

	var id int
	//tx, err := db.Begin()

	query := `insert into users (login,password) values ($1,$2) returning id` //returning login
	err = db.QueryRow(query, &login, &password).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err.Error())
		return 0, err
	}

	/*newid := u.UserExists(id)
	if newid == id {
		fmt.Println(newid)
		tx.Rollback()
	} else {
		tx.Commit()
	}*/
	return id, err

}

func (u *Repository) GetUser(user *models.User) (login, password string) {

	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}

	query := `select (login,password) from users where id=$1`
	db.QueryRow(query, login, password).Scan(user.ID)
	return login, password
}
func (u *Repository) GetUserById(id int) (user *models.User) {

	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}

	query := `select login, password from users where id=$1`
	rows := db.QueryRow(query, id)
	if err = rows.Scan(&user.ID); err != nil {
		log.Println("scan fault")
	}

	return user
}
func (u *Repository) FindIdFromLogin(user *models.User) (id int) {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}
	id = user.ID
	query := `select id from users where login=$1`
	rows := db.QueryRow(query, user.Login)
	if err = rows.Scan(&user.ID); err != nil {
		log.Println("scan fault")
	}

	return id
}

func Connect(host, port, dbname, user, password string) string {
	psqlconn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	return psqlconn
}

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", Connect("localhost", "5432", "postgres", "postgres", "root"))
	return db, err
}
