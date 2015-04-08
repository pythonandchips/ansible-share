package main

import (
	"net/url"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

type Role struct {
	host, name, tag, version string
}

func NewPushRole(tag string) Role {
	role := NewRole(tag)
	if role.version == "latest" {
		role.version = newUuid()
	}
	return role
}

func newUuid() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}

func NewRole(tag string) Role {
	indexOfTagStart := strings.LastIndex(tag, ":")
	indexOfPath := strings.IndexRune(tag, '/')
	version := newUuid()
	if indexOfPath < indexOfTagStart && indexOfTagStart != -1 {
		version = tag[indexOfTagStart+1 : len(tag)]
		tag = tag[0:indexOfTagStart]
	}
	tag = "http://" + tag

	url, err := url.Parse(tag)
	if err != nil {
		panic(err)
	}
	return Role{host: url.Host, name: url.Path, tag: tag, version: version}
}

func (serverRole Role) Url() string {
	return "http://" + serverRole.host + "/roles" + serverRole.name + "/" + serverRole.version
}
