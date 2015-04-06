package main

import "testing"

func TestCorrectParametersPulledFromTag(t *testing.T) {
	tag := "ansible.1pcdev.com/nginx:v1.1"

	role := NewRole(tag)

	url := "ansible.1pcdev.com"
	roleName := "/nginx"
	version := "v1.1"

	StringEqual(url, role.host, t)
	StringEqual(roleName, role.name, t)
	StringEqual(version, role.version, t)
}

func TestCorrectParametersPulledFromTagWithPortAndVersion(t *testing.T) {
	tag := "ansible.1pcdev.com:3030/nginx:v1.1"

	role := NewRole(tag)

	url := "ansible.1pcdev.com:3030"
	roleName := "/nginx"
	version := "v1.1"

	StringEqual(url, role.host, t)
	StringEqual(roleName, role.name, t)
	StringEqual(version, role.version, t)
}

func TestCorrectParametersPulledFromTagWithPortAndNoVersion(t *testing.T) {
	tag := "ansible.1pcdev.com:3030/nginx"

	role := NewRole(tag)

	url := "ansible.1pcdev.com:3030"
	roleName := "/nginx"

	StringEqual(url, role.host, t)
	StringEqual(roleName, role.name, t)
}

func TestCorrectParametersPulledFromTagWithLatestVersion(t *testing.T) {
	tag := "ansible.1pcdev.com:3030/nginx:latest"

	role := NewRole(tag)

	StringNotEqual("latest", role.version, t)

}

func StringEqual(expect string, actual string, t *testing.T) {
	if expect != actual {
		t.Log("expected " + expect + " to equal " + string(actual))
		t.Fail()
	}
}

func StringNotEqual(expect string, actual string, t *testing.T) {
	if expect == actual {
		t.Log("expected " + expect + " to not equal " + string(actual))
		t.Fail()
	}
}
