package access_token

import (
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
	"strings"
	"time"
)

const expirationTime = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (accessToken *AccessToken) Validate() *errors.RestErr {
	accessToken.AccessToken = strings.TrimSpace(accessToken.AccessToken)
	if accessToken.AccessToken == "" {
		return errors.BadRequestError("invalid access token id")
	}
	if accessToken.UserId <= 0 {
		return errors.BadRequestError("invalid user id")
	}
	if accessToken.ClientId <= 0 {
		return errors.BadRequestError("invalid client id")
	}
	if accessToken.Expires <= 0 {
		return errors.BadRequestError("invalid expiration")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (accessToken AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(accessToken.Expires, 0)

	return now.After(expirationTime)
}
