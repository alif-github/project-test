package in

import (
	"github.com/alif-github/project-test/model"
	"net/http"
)

type AuthUserRequest struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (input AuthUserRequest) ValidateRegisterUser() (err model.ErrorModel) {
	fileName := "AuthUserRequest.go"
	funcName := "ValidateRegisterUser"

	if input.FullName == "" {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Full Name tidak boleh kosong", fileName, funcName)
		return
	}

	if len(input.FullName) > 50 {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Panjang Full Name tidak boleh lebih dari 50", fileName, funcName)
		return
	}

	if input.Username == "" {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Username tidak boleh kosong", fileName, funcName)
		return
	}

	if len(input.Username) > 10 {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Username tidak boleh lebih dari 10", fileName, funcName)
		return
	}

	if input.Password == "" {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Password tidak boleh kosong", fileName, funcName)
		return
	}

	if len(input.Password) > 13 {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Password tidak boleh lebih dari 13", fileName, funcName)
		return
	}

	err = model.GenerateNonErrorModel()
	return
}

func (input AuthUserRequest) ValidateLoginUser() (err model.ErrorModel) {
	fileName := "AuthUserRequest.go"
	funcName := "ValidateLoginUser"

	if input.Username == "" {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Username tidak boleh kosong", fileName, funcName)
		return
	}

	if input.Password == "" {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Password tidak boleh kosong", fileName, funcName)
		return
	}

	err = model.GenerateNonErrorModel()
	return
}
