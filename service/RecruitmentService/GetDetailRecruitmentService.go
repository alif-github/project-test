package RecruitmentService

import (
	"encoding/json"
	"fmt"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/dto/out"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/util"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func (input recruitmentService) GetDetailRecruitmentService(w http.ResponseWriter, r *http.Request) (err model.ErrorModel, output out.PayloadResponseSuccess) {
	fileName := "GetDetailRecruitmentService.go"
	funcName := "GetDetailRecruitmentService"

	id, _ := mux.Vars(r)["ID"]
	if id == "" {
		err = model.GenerateErrorModel(http.StatusBadRequest, "ID tidak boleh kosong", fileName, funcName)
		return
	}

	response, errFetch := http.Get(fmt.Sprintf(`http://dev3.dansmultipro.co.id/api/recruitment/positions/%s`, id))
	if errFetch != nil {
		logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
		logModel.Code = http.StatusInternalServerError
		logModel.Message = errFetch.Error()
		util.LogError(logModel.LoggerZapFieldObject())
		err = model.GenerateErrorModel(http.StatusInternalServerError, "Kesalahan pada system, hubungi CS kami", fileName, funcName)
		return
	}

	var responseData []byte
	responseData, errFetch = ioutil.ReadAll(response.Body)
	if errFetch != nil {
		logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
		logModel.Code = http.StatusInternalServerError
		logModel.Message = errFetch.Error()
		util.LogError(logModel.LoggerZapFieldObject())
		err = model.GenerateErrorModel(http.StatusInternalServerError, "Kesalahan pada system, hubungi CS kami", fileName, funcName)
		return
	}

	var responseObject out.RecruitmentResponse
	_ = json.Unmarshal(responseData, &responseObject)

	output.Message = "Berhasil Get Data"
	output.Data = responseObject
	err = model.GenerateNonErrorModel()
	return
}
