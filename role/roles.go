package role

type Roles []Role

func (roles Roles) Exists(role Role) bool {
	exists := false
	for _, currentRole := range roles {
		if currentRole.Name == role.Name {
			exists = true
			break
		}
	}
	return exists
}

func (roles Roles) Add(role Role) Roles {
	if roles.Exists(role) {
		index := roles.IndexOf(role)
		roles[index] = role
	} else {
		roles = append(roles, role)
	}
	return roles
}

func (roles Roles) IndexOf(role Role) int {
	indexOf := -1
	for index, currentRole := range roles {
		if currentRole.Name == role.Name {
			indexOf = index
			break
		}
	}
	return indexOf
}
