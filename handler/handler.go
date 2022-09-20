package handler

import (
	"ExCloud/cache"
	_ "ExCloud/docs"
	"ExCloud/repository/postgres"
	"github.com/gin-gonic/gin"
	"github.com/gocraft/web"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (fileSrv *FileSrv) FileServerHandler() {
	http.FileServer(http.Dir("./templates/js/"))
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
}

type FileSrv struct {
	web.Router
}

type Handler struct {
	cache      *cache.Cache
	repo       *postgres.Repository
	fileServer *FileSrv
}

func NewHandler(repo *postgres.Repository, cache *cache.Cache) *Handler {
	return &Handler{repo: repo, cache: cache}
}

func InitRoutes() *gin.Engine {

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		api.POST("/register", Register)
		//api.GET("/api/register", GetData)
		api.POST("/auth", AuthenticateHandler)

		api.POST("/docs", UploadNewDocument)
		api.GET("/docs", GetDocuments)
		api.HEAD("/docs", GetDocuments)

		api.GET("/docs/<id>", GetDocumentById)
		api.HEAD("/docs/<id>", GetDocumentById)
		api.DELETE("/docs/<id>", DeleteDocumentById)

		api.GET("/auth/<token>", EndSession)
	}

	return router

	//fileserver := router.Group("file")
}
