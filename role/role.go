package role

import (
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
)

type Role struct {
	Host, Name, Tag, Version string
}

func NewPushRole(tag string) Role {
	role := NewRole(tag)
	if role.Version == "latest" {
		role.Version = newUuid()
	}
	return role
}

func newUuid() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func NewRole(tag string) Role {
	indexOfTagStart := strings.LastIndex(tag, ":")
	indexOfPath := strings.Index(tag, "/")
	host := tag[0:indexOfPath]
	name := tag[indexOfPath+1 : len(tag)]
	if indexOfTagStart > indexOfPath && indexOfTagStart != -1 {
		name = tag[indexOfPath+1 : indexOfTagStart]
	}
	version := newUuid()
	if indexOfTagStart > indexOfPath && indexOfTagStart != -1 {
		version = tag[indexOfTagStart+1 : len(tag)]
	}
	return Role{Host: host, Name: name, Tag: tag, Version: version}
}

func (role Role) AbsoluteTag() string {
	return fmt.Sprintf("%s/%s:%s", role.Host, role.Name, role.Version)
}
