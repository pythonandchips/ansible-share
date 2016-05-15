package role

import "testing"

func TestRoleParametersPulledFromTag(t *testing.T) {
	tag := "ansible.1pcdev.com/nginx:v1.1"

	role := NewRole(tag)

	host := "ansible.1pcdev.com"
	roleName := "nginx"
	version := "v1.1"

	if host != role.Host {
		t.Errorf("Expected host to be %s but got %s", host, role.Host)
	}
	if roleName != role.Name {
		t.Errorf("Expected role name to be %s but got %s", roleName, role.Name)
	}
	if version != role.Version {
		t.Errorf("Expected role version to be %s but got %s", version, role.Version)
	}
}

func TestRoleParametersPulledFromTagWithPortAndNoVersion(t *testing.T) {
	tag := "ansible.1pcdev.com:3030/nginx"

	role := NewRole(tag)

	host := "ansible.1pcdev.com:3030"
	roleName := "nginx"

	if host != role.Host {
		t.Errorf("Expected host to be %s but got %s", host, role.Host)
	}
	if roleName != role.Name {
		t.Errorf("Expected role name to be %s but go %s", roleName, role.Name)
	}
	if "" == role.Version {
		t.Errorf("Expected version to not be empty but was %s", role.Version)
	}
}

func TestRoleParametersPulledFromTagWithLatestVersion(t *testing.T) {
	tag := "ansible.1pcdev.com:3030/nginx:latest"

	role := NewRole(tag)

	if role.Version != "latest" {
		t.Errorf("Expected role version to be %s but got %s", "latest", role.Version)
	}
}

func TestRoleReplacesLatestWithUUIDWhenRoleForPush(t *testing.T) {
	tag := "ansible.1pcdev.com:3030/nginx:latest"

	role := NewPushRole(tag)

	if role.Version == "latest" {
		t.Errorf("Expected role version not to be %s but got %s", "latest", role.Version)
	}
}
