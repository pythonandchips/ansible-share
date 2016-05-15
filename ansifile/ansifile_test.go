package ansifile

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pythonandchips/ansible-share/role"
)

func TestRead(t *testing.T) {
	ioutil.WriteFile("ansifile", []byte("ansishare/base:139\nansishare/debian_base:3131"), 0755)

	roles, err := Read()

	if err != nil {
		t.Errorf("Unexpected Error %s", err)
	}
	if len(roles) != 2 {
		t.Errorf("Expected 1 role but got %d", len(roles))
	}
	if roles[1].Name != "debian_base" {
		t.Errorf("Expected role name to be %s but was %s", "debian_base", roles[0].Name)
	}
	os.Remove("ansifile")
}

func TestWrite(t *testing.T) {
	ioutil.WriteFile("ansifile", []byte("ansishare/web:3131"), 0755)

	role := role.Role{Host: "ansishare", Name: "debian_base", Version: "12342"}
	Write(role)

	roles, err := Read()
	if err != nil {
		t.Errorf("Unexpected Error %s", err)
	}
	if len(roles) != 2 {
		t.Errorf("Expected 2 roles but got %d", len(roles))
	}
}
