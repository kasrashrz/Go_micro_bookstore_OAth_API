package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/http"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/repository/db"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/repository/rest"
	access_token2 "github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	//session:= cassandra.GetSession()
	//session.Close()

	atService := access_token2.NewService(rest.NewUserRepository(), db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	//router.GET("/oath/access_token/:access_token", atHandler.GetById)
	router.Run(":4444")
}
