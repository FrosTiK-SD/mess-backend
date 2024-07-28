package utils

import (
	"slices"
	"strings"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
)

func IsInstituteEmail(email string) bool {
	at := strings.LastIndexByte(email, '@')

	if at == -1 {
		return false
	}

	domain := email[at+1:]

	if domain != "iitbhu.ac.in" && domain != "itbhu.ac.in" {
		return false
	}

	return true
}

func GetAllRoles(user *interfaces.UserPopulated) []constants.Role {
	roles := make([]constants.Role, 0, 2*len(user.Roles)+1)
	roles = append(roles, user.Roles...)
	for gidx := range user.Groups {
		for ridx := range user.Groups[gidx].Roles {
			if !slices.Contains(roles, user.Groups[gidx].Roles[ridx]) {
				roles = append(roles, user.Groups[gidx].Roles[ridx])
			}
		}
	}
	return roles
}
