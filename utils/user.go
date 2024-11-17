package utils

import (
	"strings"
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
