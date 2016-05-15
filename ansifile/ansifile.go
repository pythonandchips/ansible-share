package ansifile

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"

	"github.com/pythonandchips/ansible-share/role"
)

func Read() (role.Roles, error) {
	data, err := ioutil.ReadFile("ansifile")
	if err != nil {
		return []role.Role{}, err
	}
	buffer := bytes.NewBuffer(data)
	roles := role.Roles{}
	for {
		roleString, err := buffer.ReadString('\n')
		roleString = strings.TrimSpace(roleString)
		if roleString != "" {
			roles = roles.Add(role.NewRole(roleString))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return roles, err
		}
	}
	return roles, nil
}

func Write(role role.Role) error {
	roles, err := Read()
	if err != nil {
		return err
	}
	roles = roles.Add(role)
	ansifile := []string{}
	for _, currentRole := range roles {
		ansifile = append(ansifile, currentRole.AbsoluteTag())
	}
	data := []byte(strings.Join(ansifile, "\n"))
	ioutil.WriteFile("ansifile", data, 0755)
	return nil
}
