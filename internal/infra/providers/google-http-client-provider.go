package providers

import (
	"context"
	"net/http"

	googleoauth "github.com/henrique998/go-auth/internal/configs/google-oauth"
	"golang.org/x/oauth2"
)

type GoogleHTTPClientProvider struct {
	Token string
}

func (cp *GoogleHTTPClientProvider) Do(req *http.Request) (*http.Response, error) {
	config := googleoauth.GetGoogleOauthConfig()
	client := config.Client(context.TODO(), &oauth2.Token{
		AccessToken: cp.Token,
	})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
