package router

import (
	"net/http"
	"strings"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/ant0ine/go-json-rest/rest"
)

func middlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		dataClaims, err := tokenValidator(strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.WriteJson(map[string]string{"error": err.Error()})
			return
		}
		username := dataClaims.Get("username")
		values := r.URL.Query()
		values.Add("username", username.(string))
		r.URL.RawQuery = values.Encode()
		handler(w, r)
	}
}

func tokenValidator(tokenString string) (jwt.Claims, error) {
	token, err := jws.ParseJWT([]byte(tokenString))
	if err != nil {
		return nil, err
	}
	validator := &jwt.Validator{}
	validator.SetIssuer("foamtecintl")
	err = token.Validate([]byte(secret), crypto.SigningMethodHS256, validator)
	return token.Claims(), err
}

// SetupAPI in router
func SetupAPI(router rest.App) (api *rest.Api) {
	api = rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders := []string{
		"Accept",
		"Authorization",
		"X-Real-IP",
		"Content-Type",
		"X-Custom-Header",
		"Language",
		"Origin",
	}
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods:                allowedMethods,
		AllowedHeaders:                allowedHeaders,
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
	api.SetApp(router)
	return
}
