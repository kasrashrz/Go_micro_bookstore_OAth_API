package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/domain/access_token"
	access_token2 "github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/services/access_token"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token2.Service
}

func NewHandler(service access_token2.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

func (handler *accessTokenHandler) GetById(ctx *gin.Context) {
	accessTokenId := strings.TrimSpace(ctx.Param("access_token"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(ctx *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, accessToken)
}
