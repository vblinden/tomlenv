package tomlenv

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadSimpleString(t *testing.T) {
	testToken := "abcdef123456"

	LoadString(fmt.Sprintf("token = \"%s\"", testToken))

	if os.Getenv("token") != testToken {
		t.Fail()
	}
}

func TestLoadExtendedString(t *testing.T) {
	testUsername := "john"

	LoadString(fmt.Sprintf("[database]\n\rusername = \"%s\"", testUsername))

	if os.Getenv("database.username") != testUsername {
		t.Fail()
	}
}
