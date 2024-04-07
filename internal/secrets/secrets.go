package secrets

import "github.com/zalando/go-keyring"

var service = "mal-cli"

func Set(key string, val string) error {
	return keyring.Set(service, key, val)
}

func Get(key string) (string, error) {
	return keyring.Get(service, key)
}

func Delete(key string) error {
	return keyring.Delete(service, key)
}
