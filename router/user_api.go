package router

import (
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/foamtecintl/wpc-api/repository"
)

func userUpdateProfile(w rest.ResponseWriter, r *rest.Request) {
	body, resp := setupBody(r)
	user, err := userOnLogin(r)
	body["id"] = strconv.Itoa(user.ID)
	if err != nil {
		resp["message"] = err.Error()
		w.WriteHeader(403)
		w.WriteJson(resp)
		return
	}
	err = repository.UpdateUser(body)
	if err != nil {
		resp["message"] = err.Error()
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	resp["message"] = "success"
	w.WriteHeader(200)
	w.WriteJson(resp)
	return
}

func findUserByUsername(w rest.ResponseWriter, r *rest.Request) {
	body, resp := setupBody(r)
	_, err := userOnLogin(r)
	if err != nil {
		resp["message"] = err.Error()
		w.WriteHeader(403)
		w.WriteJson(resp)
		return
	}
	user, err := repository.FindAppUserByUsername(body["username"])
	if err != nil {
		resp["message"] = err.Error()
		w.WriteHeader(403)
		w.WriteJson(resp)
		return
	}
	if user.Username == "" {
		resp["message"] = "no data"
		w.WriteHeader(403)
		w.WriteJson(resp)
		return
	}
	resp["user"] = user
	w.WriteHeader(200)
	w.WriteJson(resp)
	return
}
