package access_token

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time constant must be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(),"Brand new access token should not be nil")
	assert.EqualValues(t, "", at.AccessToken, "New access token should not have defined access token id")
	assert.True(t, at.UserId == 0, "New access token should not have an associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expires three hours from now should not be expired")
}
