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

	fmt.Printf("Visit the URL for the auth dialog: %v", url)
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

// ?code=def502003f3a8c459899d2b2e14082d1bbcca8d1ac279e779933476519378ce1d9db03362be460a67c3202bcd5c7bd3f7867f7a0f250358d55e0abea9df83af057f6edd0222a3bf3804298975261661f6747dd1b32d7d7f826fb869d7955acf1a7e45e57b1fcbaa40680632e6fc2bac3082d0c09f334113e6f0fad4e5256a97be2f4557fd177d1c99dcc751896c0b0abb295775e631a841967784ae5b719b9d1a9bf10ddbc484da1ab55a8fbe46be5fc50fccbaeeb7069fa9bcd3433f1be2cfa3d3585ca3b4dc40be0f01f9591b785807b5e0c1f9ae7b3b53525b64f4a21c1ba9bd66367a09e2cf8b9836276b5431679a0fa6ae9b8ada088a49fa06b1a19bf5fbad176074f16fd28e8c06e6df7392af62af93f1e7139d30601e57a05c3202be9632bacf1c20f5ffb9ed7437bda3b1159e1d82c85497b81abe8fb8647fb835208153bb5792c0073eab5009610fc905b0b9d511228d2077538a87bfa009e02cecca25ffa6bfa6a54a70fdc602e56d3b76480e29b0975df54e4a9a2b2ba7bc10b19dda73cb93fe17f44b70e712a29ec56507ee95e87
// &state=tubDnKnudUtb5ISZX1S9GRfNL1yMsMPIfjUBWyF_LCQ
