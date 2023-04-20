package AuthService

import (
	"database/sql"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/dao"
	"github.com/alif-github/project-test/dto/in"
	"github.com/alif-github/project-test/dto/out"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/repository"
	"github.com/alif-github/project-test/serverconfig"
	"github.com/alif-github/project-test/service"
	"github.com/alif-github/project-test/util"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type authService struct {
	service.AbstractService
}

var AuthService = authService{}.Initiate()

func (input authService) Initiate() (output authService) {
	output.FileName = "AuthService.go"
	return
}

func (input authService) RegisterUser(response http.ResponseWriter, request *http.Request) (err model.ErrorModel, output out.PayloadResponseSuccess) {
	var (
		inputStruct *in.AuthUserRequest
		db          = serverconfig.ServerAttribute.DBConnection
	)

	_, err = input.ReadInput(request, &inputStruct)
	if err.Error != nil {
		return
	}

	err = inputStruct.ValidateRegisterUser()
	if err.Error != nil {
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(inputStruct.Password), bcrypt.DefaultCost)
	inputStruct.Password = string(hashPassword)

	_, err = dao.AuthUserDAO.RegisterNewUser(db, repository.AuthUserModel{
		FullName: sql.NullString{String: inputStruct.FullName},
		Username: sql.NullString{String: inputStruct.Username},
		Password: sql.NullString{String: inputStruct.Password},
	})
	if err.Error != nil {
		return
	}

	output = out.PayloadResponseSuccess{
		Message: "Sukses Register User Baru",
	}

	err = model.GenerateNonErrorModel()
	return
}

func (input authService) Login(response http.ResponseWriter, request *http.Request) (err model.ErrorModel, output out.PayloadResponseSuccess) {
	var (
		funcName    = "Login"
		inputStruct *in.AuthUserRequest
		db          = serverconfig.ServerAttribute.DBConnection
		userOnDB    repository.AuthUserModel
	)

	_, err = input.ReadInput(request, &inputStruct)
	if err.Error != nil {
		return
	}

	err = inputStruct.ValidateLoginUser()
	if err.Error != nil {
		return
	}

	userOnDB, err = dao.AuthUserDAO.GetUserByUsername(db, repository.AuthUserModel{Username: sql.NullString{String: inputStruct.Username}})
	if err.Error != nil {
		return
	}

	if userOnDB.ID.Int64 < 1 {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Username atau password anda salah!", input.FileName, funcName)
		return
	}

	bError := bcrypt.CompareHashAndPassword([]byte(userOnDB.Password.String), []byte(inputStruct.Password))
	if bError != nil {
		err = model.GenerateErrorModel(http.StatusBadRequest, "Username atau password anda salah!", input.FileName, funcName)
		return
	}

	//--- JWT Claims
	expiredTime := time.Now().Add(5 * time.Minute)
	claims := model.JWTClaim{
		Username: inputStruct.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-dans-multi-pro",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, tError := tokenAlgo.SignedString(model.JWT_KEY)
	if tError != nil {
		logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
		logModel.Code = http.StatusInternalServerError
		logModel.Message = tError.Error()
		util.LogError(logModel.LoggerZapFieldObject())
		err = model.GenerateErrorModel(http.StatusInternalServerError, "Kesalahan pada system, hubungi CS kami", input.FileName, funcName)
		return
	}

	//--- Set token to Cookie
	http.SetCookie(response, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	output = out.PayloadResponseSuccess{Message: "Login Berhasil"}
	err = model.GenerateNonErrorModel()
	return
}

func (input authService) Logout(response http.ResponseWriter, _ *http.Request) (err model.ErrorModel, output out.PayloadResponseSuccess) {
	http.SetCookie(response, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	output = out.PayloadResponseSuccess{Message: "Logout Berhasil"}
	err = model.GenerateNonErrorModel()
	return
}
