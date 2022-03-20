package db

import (
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/domain/access_token"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {

}

func (repo *dbRepository)GetById(string) (*access_token.AccessToken, *errors.RestErr){
	return nil, errors.InternalServerError("database connection not implemented yet!")

}