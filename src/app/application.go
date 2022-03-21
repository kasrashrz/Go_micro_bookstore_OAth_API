package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/clients/cassandra"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/domain/access_token"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/http"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr !=nil {
		panic(dbErr)
	}
	session.Close()

	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oath/access_token/:access_token", atHandler.GetById)
	router.Run(":4444")
}
