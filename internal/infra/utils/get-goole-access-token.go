package utils

import (
	"context"

	googleoauth "github.com/henrique998/go-auth/internal/configs/google-oauth"
)

func GetGoogleAccessToken(code string) (accessToken string, err error) {
	config := googleoauth.GetGoogleOauthConfig()

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, err
}
