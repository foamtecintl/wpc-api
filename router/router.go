package router

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
)

var secret string

//NewRouter is create router
func NewRouter(secretKey string) rest.App {
	secret = secretKey
	router, err := rest.MakeRouter(
		rest.Post("/api/login", login),
		rest.Post("/api/validateusername", validateUsername),
		rest.Post("/api/validateemployeeid", validateEmployeeID),
		rest.Post("/api/register", register),
		rest.Post("/api/updateprofile", middlewareFunc(userUpdateProfile)),
		rest.Post("/api/finduserbyusername", middlewareFunc(findUserByUsername)),
	)
	if err != nil {
		log.Fatal(err)
	}
	return router
}
