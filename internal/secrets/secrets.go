package secrets

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

var service = "mal-cli"

func Set(key string, val string) {
	err := keyring.Set(service, key, val)
	if err != nil {
		fmt.Println(err)
	}
}

func Get(key string) string {
	secret, err := keyring.Get(service, key)
	if err != nil {
		fmt.Println(err)
	}

	return secret
}

func Delete(key string) {
	err := keyring.Delete(service, key)
	if err != nil {
		fmt.Println(err)
	}
}
