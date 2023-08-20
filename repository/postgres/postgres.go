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
	*sql.DB
	cache cache.Cache
	log   log.Logger
	UserPostgres
}

type UserPostgres interface {
	CreateUser(login, password string) error
	AuthenticateStorage(login, password string) error
	GetUserById(id string) error
}

func (u *Repository) GetAllUsers() (*models.User, error) {
	var us models.User
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}
	query := `select * from users`
	err = db.QueryRow(query).Scan(&us.ID, &us.Login, &us.Password)
	if err != nil {
		log.Println(err.Error())
	}
	return &us, err
	/*rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []models.User{}

	for rows.Next() {
		p := models.User{}
		err := rows.Scan(&p.ID, &p.Login, &p.Password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	return products, err*/
}

func (u *Repository) CreateUser(login, password string) (int, error) {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}
	var id int
	query := `insert into users (login,password) values ($1,$2) returning id` //returning login
	err = db.QueryRow(query, &login, &password).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err.Error())
		return 0, err
	}
	return id, err
}

func (u *Repository) AuthenticateStorage(login, password string) (string, error) {

	user := models.User{}
	query := `SELECT login,password FROM users WHERE login=$1 AND password=$2`
	err := u.QueryRow(query, &login, &password).Scan(&user.Login, &user.Password)
	if err != nil {
		return "", err
	}

	return "", err
}

/*func (u *Repository) TokenExists(login string) bool {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}

	var exists bool
	//query := `SELECT EXISTS (SELECT id FROM users WHERE id = $1)`
	qry := `SELECT EXISTS (SELECT token FROM users WHERE login = $1);`
	db.QueryRow(qry, login).Scan(&exists)

	return exists
}*/

func (u *Repository) LoginExists(login string) bool {
	db, err := OpenConn()
	if err != nil {
		log.Println(err.Error())
	}
	var exists bool
	qry := `SELECT EXISTS (SELECT login FROM users WHERE login =$1);`
	db.QueryRow(qry, login).Scan(&exists)
	return exists
}

func (u *Repository) UserExists(login, password string) bool {
	var exists bool
	qry := `SELECT EXISTS (SELECT login,password FROM users WHERE login =$1 AND password=$2);`
	u.QueryRow(qry, &login, &password).Scan(&exists)
	return exists
}

func (u *Repository) GetUser(user *models.User) (login, password string) {
	query := `select (login,password) from users where id=$1`
	u.QueryRow(query, login, password).Scan(user.ID)
	return login, password
}

func (u *Repository) GetUserById(id int) (user *models.User) {
	query := `select login, password from users where id=$1`
	rows := u.QueryRow(query, id)
	if err := rows.Scan(&user.ID); err != nil {
		log.Println("scan fault")
	}
	return user
}

func (u *Repository) FindIdFromLogin(user *models.User) (id int) {
	id = user.ID
	query := `select id from users where login=$1`
	rows := u.QueryRow(query, user.Login)
	if err := rows.Scan(&user.ID); err != nil {
		log.Println("scan fault")
	}
	return id
}

func Connect(host, port, dbname, user, password string) string {
	psqlconn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	return psqlconn
}

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", Connect("localhost", "5432", "postgres", "postgres", "postgres"))
	return db, err
}
