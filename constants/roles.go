package constants

type Role = string

const (
	USER         Role = "USER"
	CARETAKER    Role = "CARETAKER"
	WARDEN       Role = "WARDEN"
	ADMIN_WARDEN Role = "ADMIN_WARDEN"
	ADMIN        Role = "ADMIN"
)
