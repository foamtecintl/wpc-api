package router

import (
	"errors"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/foamtecintl/wpc-api/repository"
	"golang.org/x/crypto/bcrypt"
)

func validateUsername(w rest.ResponseWriter, r *rest.Request) {
	body, resp := setupBody(r)
	err := checkUsername(body["username"])
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

func validateEmployeeID(w rest.ResponseWriter, r *rest.Request) {
	body, resp := setupBody(r)
	err := checkEmployeeID(body["employeeId"])
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

func login(w rest.ResponseWriter, r *rest.Request) {
	body, resp := setupBody(r)
	user, err := repository.FindAppUserByUsername(body["username"])
	token, err := validateUserPass(user, body["password"])
	if err != nil {
		resp["message"] = err.Error()
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	resp["token"] = token
	resp["username"] = user.Username
	resp["role"] = user.RoleName
	resp["status"] = user.Status
	w.WriteHeader(200)
	w.WriteJson(resp)
	return
}

func validateUserPass(user repository.User, password string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid user or pass")
	}
	claims := jws.Claims{}
	claims.Set("username", user.Username)
	claims.SetIssuer("foamtecintl")
	now := time.Now()
	claims.SetIssuedAt(now)
	claims.SetExpiration(now.AddDate(0, 0, 3))
	tokenStruct := jws.NewJWT(claims, crypto.SigningMethodHS256)
	serialized, err := tokenStruct.Serialize([]byte(secret))
	if err != nil {
		return "", err
	}
	return string(serialized), nil
}

func checkUsername(username string) error {
	user, err := repository.FindAppUserByUsername(username)
	if err != nil {
		return err
	}
	if user.Username != "" {
		return errors.New("username duplicate")
	}
	return nil
}

func checkEmployeeID(employeeID string) error {
	user, err := repository.FindAppUserByEmployeeID(employeeID)
	if err != nil {
		return err
	}
	if user.EmployeeID != "" {
		return errors.New("employee id duplicate")
	}
	return nil
}
