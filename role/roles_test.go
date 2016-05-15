package role

import "testing"

func TestRolesAddNewRole(t *testing.T) {
	roles := Roles{}

	role := Role{Name: "debian_base"}

	roles = roles.Add(role)

	if len(roles) != 1 {
		t.Errorf("Expected roles to have %d roles but got %d", 1, len(roles))
	}
}
