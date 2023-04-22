package RecruitmentService

import (
	"encoding/json"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/dto/out"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/util"
	"github.com/gorilla/mux"
	"net/http"
)

func (input recruitmentService) GetDetailRecruitmentService(w http.ResponseWriter, r *http.Request) (err model.ErrorModel, output out.PayloadResponseSuccess) {
	var (
		fileName       = "GetDetailRecruitmentService.go"
		funcName       = "GetDetailRecruitmentService"
		errFetch       error
		responseData   []byte
		responseObject out.RecruitmentResponse
	)

	id, _ := mux.Vars(r)["ID"]
	if id == "" {
		err = model.GenerateErrorModel(http.StatusBadRequest, "ID tidak boleh kosong", fileName, funcName)
		return
	}

	defer func() {
		if errFetch != nil {
			logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
			logModel.Code = http.StatusInternalServerError
			logModel.Message = errFetch.Error()
			util.LogError(logModel.LoggerZapFieldObject())
		}
	}()

	responseData, errFetch, err = input.fetchDataHitAPI(config.ApplicationConfiguration.GetExternalAPI().Url + config.ApplicationConfiguration.GetExternalAPI().Path.View + "/" + id)
	if errFetch != nil {
		return
	}

	_ = json.Unmarshal(responseData, &responseObject)

	output.Message = "Berhasil Get Data"
	output.Data = responseObject
	err = model.GenerateNonErrorModel()
	return
}
