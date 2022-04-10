package access_token

import (
	"fmt"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/crypto_utils"
	"github.com/kasrashrz/Go_micro_bookstore_OAth_API/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credential"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	// For password grant type
	Username string `json:"email"`
	Password string `json:"password"`
	// For user_credential grand_type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (accessToken *AccessTokenRequest) Validate() *errors.RestErr {
	switch accessToken.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.BadRequestError("invalid grant_type parameter")
	}

	//TODO: Validate parameters for each grant type
	return nil
}

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

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (accessToken AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(accessToken.Expires, 0)

	return now.After(expirationTime)
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
