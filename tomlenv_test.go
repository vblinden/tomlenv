package tomlenv

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadSimpleString(t *testing.T) {
	testToken := "abcdef123456"

	err := LoadString(fmt.Sprintf("token = \"%s\"", testToken))
	if err != nil {
		t.Fail()
	}

	if os.Getenv("token") != testToken {
		t.Fail()
	}
}

func TestLoadExtendedString(t *testing.T) {
	testUsername := "john"

	err := LoadString(fmt.Sprintf("[database]\n\rusername = \"%s\"", testUsername))
	if err != nil {
		t.Fail()
	}

	if os.Getenv("database.username") != testUsername {
		t.Fail()
	}
}
