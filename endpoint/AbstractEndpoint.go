package endpoint

import (
	"encoding/json"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/dto/out"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/util"
	"github.com/google/uuid"
	"net/http"
)

type AbstractEndpoint struct {
	FileName string
}

func (input AbstractEndpoint) ServeJWTTokenValidationEndpoint(response http.ResponseWriter, request *http.Request, serveFunction func(http.ResponseWriter, *http.Request) (model.ErrorModel, out.PayloadResponseSuccess)) {
	serveJWTTokenValidationEndpoint(response, request, serveFunction)
}

func serveJWTTokenValidationEndpoint(response http.ResponseWriter, request *http.Request, serveFunction func(http.ResponseWriter, *http.Request) (model.ErrorModel, out.PayloadResponseSuccess)) {
	var (
		err     model.ErrorModel
		payload out.PayloadResponseSuccess
	)

	defer func() {
		if err.Error != nil {
			failedProcess(err, response)
		} else {
			successProcess(response, payload)
		}
	}()

	err, payload = serveFunction(response, request)
	if err.Error != nil {
		return
	}
}

func successProcess(response http.ResponseWriter, payload out.PayloadResponseSuccess) {
	var (
		requestID string
		logModel  model.LoggerModel
	)

	requestID = uuid.New().String()
	logModel = model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
	logModel.RequestID = requestID
	logModel.Code = 200

	util.LogInfo(logModel.LoggerZapFieldObject())

	payload.Success = true
	payload.RequestID = requestID
	result, _ := json.Marshal(payload)
	response.WriteHeader(http.StatusOK)
	response.Header().Add("Content-Type", "application/json")
	_, _ = response.Write(result)
}

func failedProcess(err model.ErrorModel, response http.ResponseWriter) {
	var (
		payload   out.PayloadResponseFailed
		requestID string
		logModel  model.LoggerModel
	)

	requestID = uuid.New().String()
	logModel = model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
	logModel.RequestID = requestID
	logModel.Code = err.Code
	logModel.Message = err.Error.Error()

	util.LogError(logModel.LoggerZapFieldObject())

	if err.Error != nil {
		payload = out.PayloadResponseFailed{
			Success:      false,
			RequestID:    requestID,
			ErrorCode:    int64(err.Code),
			ErrorMessage: err.Error.Error(),
		}
	}

	result, _ := json.Marshal(payload)
	if err.Code == http.StatusInternalServerError {
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.WriteHeader(http.StatusBadRequest)
	}

	response.Header().Add("Content-Type", "application/json")
	_, _ = response.Write(result)
}
