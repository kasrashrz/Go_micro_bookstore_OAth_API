package access_token

import (
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr

}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}

}

func (service *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.BadRequestError("invalid access token id")
	}
	accessToken, err := service.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil

}

func (service *service) Create(token AccessToken) *errors.RestErr{
	if err := token.Validate(); err != nil{
		return err
	}
	return service.repository.Create(token)
}

func (service *service) UpdateExpirationTime(token AccessToken) *errors.RestErr{
	if err := token.Validate(); err != nil{
		return err
	}
	return service.repository.UpdateExpirationTime(token)
}
