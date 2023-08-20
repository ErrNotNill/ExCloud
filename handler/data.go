package handler

import (
	"ExCloud/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

const AdminToken = "test"

func Serv(c *gin.Context) {
	c.Writer.Write([]byte("SD"))
	//w.Write([]byte("SD"))
	http.FileServer(http.Dir("./templates/css/"))
}

/*func GetDataWIthTemplates(c *gin.Context) {
//todo file server handler
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

// ####### SWAGGER ##########
// @Summary Register
// @tags sign-up
// @Description register
// @Id login
// @Accept json
// @Produce json

// Register @Schemes //edited
// @Description user registration
// @Tags Register
// @Accept json
// @Produce json
// @Success 200 {"response":{"login":"string","password":"string"}} string "token"
// @Router /api/register [post]
func (h *Handler) Register(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	login := c.PostForm("login")
	password := c.PostForm("password")

	if c.Request.Method != http.MethodPost {
		c.Writer.WriteHeader(http.StatusBadRequest)
	} else {

		if CheckLoginCase(login) && CheckPasswordCase(password) == true {
			fmt.Println("ok")
			if h.repo.LoginExists(login) == true {
				fmt.Println("login exist")
				c.Writer.WriteHeader(http.StatusBadRequest)
				c.Writer.Write([]byte("This user already exists"))
			} else {
				_, err := h.repo.CreateUser(login, password)

				if err != nil {
					c.Writer.WriteHeader(http.StatusBadRequest)
					c.Writer.Write([]byte("This user already exists"))
				} else {
					c.JSON(http.StatusOK, map[string]string{
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

// @Summary AuthenticateHandler
// @tags sign-in
// @Description login
// @Id login
// @Accept json
// @Produce json

// AuthenticateHandler @Schemes //edited
// @Description user authenticate
// @Tags AuthenticateHandler
// @Accept json
// @Produce json
// @Success 200 {"request":{"login":string,"password":string} "response":{"token":string}
// @Router /api/auth [post]
func (h *Handler) AuthenticateHandler(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.Writer.WriteHeader(http.StatusBadRequest)
	} else {

		login := c.PostForm("login")
		password := c.PostForm("password")
		//Middleware(c)
		if h.repo.UserExists(login, password) == true {
			token, err := h.repo.AuthenticateStorage(login, password)

			//	userToken := models.User{Token: token}
			if err != nil {
				log.Fatalln("token generate err data/126")
				return
			}

			c.JSON(http.StatusOK, map[string]string{
				"token": token,
			})
			fmt.Println(token)

		} else {
			c.JSON(http.StatusBadRequest, nil)
		}
		/*err = repo.AuthenticateStorage(login, password)
		if err != nil {
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}*/

		//GenerateToken() the token must be assigned to the user
	}
}

//LoadFile todo repair FileLoader
/*func (h *Handler) LoadFile(filename string, targetUrl string) error {
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
}*/

// UploadNewDocument todo fix this handler
func (h *Handler) UploadNewDocument(c *gin.Context) {

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

func (h *Handler) GetDocuments(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.Writer.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		c.Writer.Write([]byte("HI"))
	}
}
func (h *Handler) GetAllUsers(c *gin.Context) {
	//user := models.User{}

	users, err := h.repo.GetAllUsers()
	if err != nil {
		fmt.Println("WTF?")
	}
	msg, err := json.Marshal(users)
	var data []models.User
	err = json.Unmarshal(msg, &data)
	/*var orders = make([]string, 0)
	ja, _ := json.Marshal(users)
	json.Unmarshal(ja, &user)
	ordersFromHabr := string(ja)
	orders = append(orders, ordersFromHabr)
	fmt.Println(orders)*/
	//userY := &models.User{Login: users.Login, Password: users.Password}
	//users := models.User{}
	if c.Request.Method != http.MethodGet {
		c.Writer.WriteHeader(http.StatusMethodNotAllowed)
	} else {

		//ja, _ := json.MarshalIndent(users, "", "")
		c.Writer.Write(msg)

	}
}

func (h *Handler) GetDocumentById(c *gin.Context) {
	c.Param("<id>")
	/*id, err := strconv.Atoi(r.URL.Query().Get("<id>"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Отображение ID %d...", id)*/
}

func (h *Handler) DeleteDocumentById(c *gin.Context) {
	//FileServer
}

func (h *Handler) EndSession(c *gin.Context) {
	c.Request.URL.Query().Get("id")
}
