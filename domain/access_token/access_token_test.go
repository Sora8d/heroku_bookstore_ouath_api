package access_token

import "testing"

func TestAccessTokenIsExpired(t *testing.T) {
	at := GetNewAccessToken(100, false)
	if at.IsExpired() {
		t.Error("brand new access token should not be expired")
	}
}
