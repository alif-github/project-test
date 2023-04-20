package endpoint

import (
	"github.com/alif-github/project-test/service/AuthService"
	"net/http"
)

type authEndpoint struct {
	AbstractEndpoint
}

var AuthEndpoint = authEndpoint{}.Initiate()

func (input authEndpoint) Initiate() (output authEndpoint) {
	input.FileName = "AuthEndpoint.go"
	return
}

func (input authEndpoint) RegisterUser(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		input.ServeJWTTokenValidationEndpoint(response, request, AuthService.AuthService.RegisterUser)
	}
}

func (input authEndpoint) Login(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		input.ServeJWTTokenValidationEndpoint(response, request, AuthService.AuthService.Login)
	}
}

func (input authEndpoint) Logout(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		input.ServeJWTTokenValidationEndpoint(response, request, AuthService.AuthService.Logout)
	}
}
