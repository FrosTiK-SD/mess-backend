package handler

import (
	"github.com/FrosTiK-SD/mess-backend/models"
	mongik "github.com/FrosTiK-SD/mongik/models"
	"github.com/lestrrat-go/jwx/jwk"
)

type Handler struct {
	MongikClient *mongik.Mongik
	JwkSet       *jwk.Set
	Session      *Session
}

type Session struct {
	Error error
	User  models.User
}
