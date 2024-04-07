package secrets

import (
	"testing"

	"github.com/zalando/go-keyring"
)

// TestSetGet calls secrets.Set with a key and a value and expects no error.
// Then it calls secrets.Get with the same key and expects the same value.
func TestSetGet(t *testing.T) {
	keyring.MockInit()
	err := Set("testKey", "testValue")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	val, err := Get("testKey")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if val != "testValue" {
		t.Errorf("Expected 'testValue', got %s", val)
	}
}

// TestDelete calls secrets.Set with a key and a value, then calls secrets.Delete
// with the same key and expects no error. Finally, it calls secrets.Get with the
// same key and expects an empty string.
func TestDelete(t *testing.T) {
	keyring.MockInit()
	err := Set("testKey", "testValue")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	err = Delete("testKey")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	val, err := Get("testKey")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if val != "" {
		t.Errorf("Expected '', got %s", val)
	}
}

// TestSetError calls secrets.MockInitWithError with an error, then calls secrets.Set
// with a key and a value and expects the error to be printed.
func TestSetError(t *testing.T) {
	keyring.MockInitWithError(keyring.ErrNotFound)
	err := Set("testKey", "testValue")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
