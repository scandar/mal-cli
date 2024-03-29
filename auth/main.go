package auth

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

var conf = &oauth2.Config{
	ClientID: "2ce6b399a4808b61741e0e121ff24661",
	Endpoint: oauth2.Endpoint{
		TokenURL: "https://myanimelist.net/v1/oauth2/token",
		AuthURL:  "https://myanimelist.net/v1/oauth2/authorize",
	},
}

func GenerateRedirectURI() string {
	verifier := oauth2.GenerateVerifier()
	challenge := oauth2.S256ChallengeFromVerifier(verifier)
	codeChallenge := oauth2.SetAuthURLParam("code_challenge", challenge)
	url := conf.AuthCodeURL(verifier, codeChallenge)
	fmt.Println("If you're not redirected automatically, please copy and paste the following URL into your browser.")
	fmt.Println(url)
	return url
}

func Exchange(verifier string, code string, opts ...oauth2.AuthCodeOption) *oauth2.Token {
	challenge := oauth2.S256ChallengeFromVerifier(verifier)
	ctx := context.Background()
	opts = append(opts, oauth2.SetAuthURLParam("code_verifier", challenge))
	token, err := conf.Exchange(ctx, code, opts...)
	if err != nil {
		fmt.Printf("Failed to exchange token: %v", err.Error())
		panic(err)
	}
	return token
}
