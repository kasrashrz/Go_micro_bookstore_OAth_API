package access_token

import (
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/domain/access_token"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/repository/db"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/repository/rest"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUserRepository
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RestUserRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (service *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.BadRequestError("invalid access token id")
	}
	accessToken, err := service.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil

}

func (service *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	//if err := request.Validate(); err != nil {
	//	return nil, err
	//}
	////TODO: Support both grant types: client_credentials and password
	//// Authenticate the user against the Users API:
	//user, err := service.restUsersRepo.LoginUser(request.Username, request.Password)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Generate a new access token:
	at := access_token.GetNewAccessToken(1)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := service.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (service *service) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return service.UpdateExpirationTime(token)
}
