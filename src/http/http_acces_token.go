package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/domain/access_token"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

func (handler *accessTokenHandler) GetById(ctx *gin.Context){
	accessTokenId := strings.TrimSpace(ctx.Param("access_token"))

	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, accessToken)
}
