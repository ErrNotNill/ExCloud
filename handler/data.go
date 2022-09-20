package handler

import (
	"ExCloud/repository/postgres"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const AdminToken = "test"

func Serv(c *gin.Context) {
	c.Writer.Write([]byte("SD"))
	//w.Write([]byte("SD"))
	http.FileServer(http.Dir("./templates/css/"))
}

/*func GetData(c *gin.Context) {
	c.Sta
	return func(w http.ResponseWriter, r *http.Request) {

	}

	http.FileServer(http.Dir("/templates/css/"))
	pages := []string{
		"./templates/html/register/home.page.tmpl",
		"./templates/html/register/base.layout.tmpl",
	}
	//tmpl := template.Must(template.ParseFiles(pages...))
	tmpl, _ := template.ParseFiles(pages...)

	err := tmpl.Execute(nil, Register) //instead of nil must be w (ResponseWriter)
	if err != nil {
		log.Println(err.Error())
	}

}*/

// @Summary Register
// @tags sign-up
// @Description login
// @Id login
// @Accept json
// @Produce json

// @Schemes
// @Description user registration
// @Tags Register
// @Accept json
// @Produce json
// @Success 200 {"response":{"login":"string"}} string "token"
// @Router /api/register [post]
func Register(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	repo := postgres.Repository{}
	login := c.PostForm("login")
	password := c.PostForm("password")

	if c.Request.Method != http.MethodPost {
		c.Writer.WriteHeader(http.StatusBadRequest)
	} else {
		if CheckLoginCase(login) == true && CheckPasswordCase(password) == true {
			if repo.LoginExists(login) == true {
				c.Writer.WriteHeader(http.StatusBadRequest)
				c.Writer.Write([]byte("This user already exists"))
			} else {
				_, err := repo.CreateUser(login, password)
				if err != nil {
					c.Writer.WriteHeader(http.StatusBadRequest)
					c.Writer.Write([]byte("This user already exists"))
				} else {
					c.JSON(http.StatusOK, map[string]interface{}{
						"login": login,
					})
					c.Writer.WriteHeader(http.StatusOK)
					fmt.Printf("user created with login: %v\n", login)
				}
			}

		} else {
			c.Writer.Write([]byte("Login or password form incorrect"))
			c.Writer.WriteHeader(http.StatusBadRequest)
		}
	}
}

func AuthenticateHandler(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.Writer.WriteHeader(http.StatusBadRequest)
	} else {
		repo := postgres.Repository{}

		login := c.PostForm("login")
		password := c.PostForm("password")

		if repo.UserExists(login, password) == true {
			c.Writer.Write([]byte("Welcome"))
			token, err := GenerateToken(password, time.Second*2)
			if err != nil {
				log.Fatalln("token generate err data/126")
				return
			}
			c.JSON(http.StatusOK, map[string]string{
				"token": token,
			})
			fmt.Println(token)
		}
		/*err = repo.AuthenticateStorage(login, password)
		if err != nil {
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}*/

		//GenerateToken() the token must be assigned to the user
	}
}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func UploadNewDocument(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "multipart/form-data")
	if c.Request.Method != "POST" {
		fmt.Fprintf(c.Writer, "Error is %v", http.StatusMethodNotAllowed)
	} else {
		/*target_url := "http://localhost:8000/api/docs"
		filename := "./astaxie.pdf"
		postFile(filename, target_url)*/
		c.Writer.Write([]byte("HELLO"))

		fmt.Println("method:", c.Request.Method)

		/*crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))*/
		//fmt.Println("HELLO", token)

		/*t, _ := template.ParseFiles("./upload.html")
		err := t.Execute(c.Writer, token)
		if err != nil {
			log.Println("html parse error")
		}*/

		c.Request.ParseMultipartForm(32 << 20)
		file, handler, err := c.Request.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(c.Writer, "%v", handler.Header)
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func GetDocuments(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.Writer.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		c.Writer.Write([]byte("HI"))
	}
}

func GetDocumentById(c *gin.Context) {
	/*id, err := strconv.Atoi(r.URL.Query().Get("<id>"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Отображение ID %d...", id)*/
}
func DeleteDocumentById(c *gin.Context) {

}
func EndSession(c *gin.Context) {
	c.Request.URL.Query().Get("id")
}

func DataResponse(w http.ResponseWriter, r *http.Request) {
}
func SendData(w http.ResponseWriter, r *http.Request) {
}
