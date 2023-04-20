package router

import (
	"encoding/json"
	"errors"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/dto/out"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"net/http"
)

func Middleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		w.Header().Set("Content-Type", "application/json")
		defer func() {
			if err := recover(); err != nil {
				logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
				logModel.Code = 500
				logModel.Message = errors.New(err.(string)).Error()
				util.LogError(logModel.LoggerZapFieldObject())
			}
		}()
		nextHandler.ServeHTTP(w, r)
	})
}

func JWTMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		w.Header().Set("Content-Type", "application/json")
		var (
			payloadErr out.PayloadResponseFailed
			errModel   model.ErrorModel
			fileName   = "Middleware.go"
			funcName   = "JWTMiddleware"
			requestID  = uuid.New().String()
			token      *jwt.Token
		)

		defer func() {
			if err := recover(); err != nil {
				logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
				logModel.Code = http.StatusInternalServerError
				logModel.Message = errors.New(err.(string)).Error()
				util.LogError(logModel.LoggerZapFieldObject())
			}

			if errModel.Error != nil {
				payloadErrJSON, _ := json.Marshal(payloadErr)
				_, _ = w.Write(payloadErrJSON)
				logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
				logModel.Code = int(payloadErr.ErrorCode)
				logModel.Message = errModel.Error.Error()
				util.LogError(logModel.LoggerZapFieldObject())
			}
		}()

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				errModel = model.GenerateErrorModel(http.StatusUnauthorized, "Unauthorized", fileName, funcName)
				payloadErr = out.PayloadResponseFailed{
					Success:      false,
					RequestID:    requestID,
					ErrorCode:    int64(errModel.Code),
					ErrorMessage: errModel.Error.Error(),
				}
				return
			}
		}

		tokenString := c.Value
		claims := &model.JWTClaim{}
		token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				errModel = model.GenerateErrorModel(http.StatusUnauthorized, "Unauthorized", fileName, funcName)
				payloadErr = out.PayloadResponseFailed{
					Success:      false,
					RequestID:    requestID,
					ErrorCode:    int64(errModel.Code),
					ErrorMessage: errModel.Error.Error(),
				}
			case jwt.ValidationErrorExpired:
				errModel = model.GenerateErrorModel(http.StatusUnauthorized, "Unauthorized, Token Expired!", fileName, funcName)
				payloadErr = out.PayloadResponseFailed{
					Success:      false,
					RequestID:    requestID,
					ErrorCode:    int64(errModel.Code),
					ErrorMessage: errModel.Error.Error(),
				}
			default:
				errModel = model.GenerateErrorModel(http.StatusUnauthorized, "Unauthorized", fileName, funcName)
				payloadErr = out.PayloadResponseFailed{
					Success:      false,
					RequestID:    requestID,
					ErrorCode:    int64(errModel.Code),
					ErrorMessage: errModel.Error.Error(),
				}
			}
			return
		}

		if !token.Valid {
			errModel = model.GenerateErrorModel(http.StatusUnauthorized, "Unauthorized", fileName, funcName)
			payloadErr = out.PayloadResponseFailed{
				Success:      false,
				RequestID:    requestID,
				ErrorCode:    int64(errModel.Code),
				ErrorMessage: errModel.Error.Error(),
			}
			return
		}

		nextHandler.ServeHTTP(w, r)
	})
}
