package db

import (
	"github.com/gocql/gocql"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/clients/cassandra"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/domain/access_token"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (repo *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).
		Scan(&result.AccessToken,
			&result.UserId,
			&result.ClientId,
			&result.Expires,
		); err != nil {
		if err == gocql.ErrNotFound{
			return nil, errors.NotFoundError("no access token found")
			
		}
		return nil, errors.InternalServerError(err.Error())
	}
	return &result, nil

}

func (repo *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {



	return nil
}