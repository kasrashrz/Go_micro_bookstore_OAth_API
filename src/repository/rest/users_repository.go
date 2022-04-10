package rest

import (
	"encoding/json"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/domain/user"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*user.User, *errors.RestErr)
}

type usersRepository struct{}

func NewUserRepository() RestUserRepository {
	return &usersRepository{}
}

func (repository *usersRepository) LoginUser(email string, password string) (*user.User, *errors.RestErr) {
	request := user.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("invalid rest client response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)

		if err != nil {
			return nil, errors.InternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user user.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.InternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}
