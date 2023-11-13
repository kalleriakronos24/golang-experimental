package userRole

type Role struct {
	index int
	role  string
}

type RoleEnum struct {
	roles []Role
}

func (enum RoleEnum) RoleName(roleName string) string {
	for _, item := range enum.roles {
		if item.role == roleName {
			return item.role
		}
	}
	return "Role not found."
}

func (enum RoleEnum) RoleIndex(findName string) int {
	for idx, item := range enum.roles {
		if findName == item.role {
			return idx
		}
	}
	return -1
}

var (
	UserRoles = RoleEnum{[]Role{{index: 1, role: "superadmin"}}}
)
