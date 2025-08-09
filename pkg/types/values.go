package types

type ROLES struct {
	ADMIN    string
	CUSTOMER string
	SUPER    string
}

func ROLE() ROLES {

	var roles ROLES = ROLES{
		ADMIN:    "admin",
		CUSTOMER: "customer",
		SUPER:    "super",
	}
	return roles
}
