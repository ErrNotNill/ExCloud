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
	handler := new(Handler)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		api.POST("/register", handler.Register)
		//api.GET("/api/register", GetData)
		api.POST("/auth", handler.AuthenticateHandler)

		api.POST("/docs", handler.UploadNewDocument)
		api.GET("/docs", handler.GetDocuments)
		api.HEAD("/docs", handler.GetDocuments)

		api.GET("/docs/<id>", handler.GetDocumentById)
		api.HEAD("/docs/<id>", handler.GetDocumentById)
		api.DELETE("/docs/<id>", handler.DeleteDocumentById)

		api.GET("/auth/<token>", handler.EndSession)
	}

	return router

	//fileserver := router.Group("file")
}
