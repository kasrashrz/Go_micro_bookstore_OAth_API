package db

import (
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/clients/cassandra"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/domain/access_token"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
)

const (
	queryGetAccessToken       = "SELECT access_tokens, users_id, client_id, expires FROM access_tokens WHERE access_tokens=?;"
	queryInsertAccessToken    = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpirationTime = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (repo *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var results access_token.AccessToken
	statement, err := cassandra.Client.Prepare(queryGetAccessToken)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(id)

	if readErr := result.Scan(
		&results.AccessToken,
		&results.UserId,
		&results.ClientId,
		&results.Expires); readErr != nil {
		return nil, errors.NotFoundError("user not found")
	}
	return &results, nil
}

func (repo *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	//if err := cassandra.GetSession().Query(queryInsertAccessToken,
	//	token.AccessToken,
	//	token.UserId,
	//	token.ClientId,
	//	token.Expires).
	//	Exec(); err != nil {
	//	return errors.InternalServerError(err.Error())
	//}
	return nil
}

func (repo *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {

	//if err := cassandra.GetSession().Query(queryUpdateExpirationTime,
	//	token.Expires,
	//	token.AccessToken).
	//	Exec(); err != nil {
	//	return errors.InternalServerError(err.Error())
	//}
	return nil
}
