package router

import (
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/foamtecintl/wpc-api/repository"
	"github.com/tomasen/realip"
)

func setupBody(r *rest.Request) (map[string]string, map[string]interface{}) {
	body := map[string]string{}
	body["clientIP"] = realip.FromRequest(r.Request)
	resp := map[string]interface{}{}
	err := r.DecodeJsonPayload(&body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body, resp
}

func userOnLogin(r *rest.Request) (repository.User, error) {
	autenUsername := r.URL.Query().Get("username")
	if len(autenUsername) <= 0 {
		return repository.User{}, fmt.Errorf("No data in parameter")
	}
	return repository.FindAppUserByUsername(autenUsername)
}
