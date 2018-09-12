package router

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/foamtecintl/wpc-api/repository"
)

func register(w rest.ResponseWriter, r *rest.Request) {
	body, resp := setupBody(r)
	err := repository.UserRegister(body)
	if err != nil {
		resp["message"] = err.Error()
		w.WriteHeader(403)
		w.WriteJson(resp)
		return
	}
	resp["message"] = "success"
	w.WriteHeader(200)
	w.WriteJson(resp)
	return
}
