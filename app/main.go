package main

import (
	//	_ "ExCloud/app/docs"
	"ExCloud/handler"
	"ExCloud/models"
	"ExCloud/models/user/register"
	"encoding/json"
	"fmt"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/http-swagger"
	"html/template"
	"log"
	"net/http"
)

// @title App API
// @version 1.0
// @description Application API

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	/*r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"), //The url pointing to API definition
	))

	http.ListenAndServe(":1323", r)*/

	err := http.ListenAndServe(":8000", handler.InitRoutes())
	if err != nil {
		log.Fatalln("srv not started")
	}
}

/*mux := http.NewServeMux()
mux.HandleFunc("/reg", regHandler)
mux.HandleFunc("/auth", authHandler)*/

func regHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
	}

	user := new(models.User)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Writer()
	}

	data := register.UserRegistration{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	fmt.Println("1")
	ts, err := template.ParseFiles("/templates/html/register/register.html")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("2")
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
	}
	//ts.Execute(w, &users)

	fmt.Println(user)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
	}

	w.Write([]byte("World"))
}
