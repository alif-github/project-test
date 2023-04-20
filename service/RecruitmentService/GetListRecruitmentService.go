package RecruitmentService

import (
	"encoding/json"
	"fmt"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/dto/in"
	"github.com/alif-github/project-test/dto/out"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/util"
	"net/http"
	"net/url"
	"strconv"
)

func (input recruitmentService) GetListRecruitmentService(w http.ResponseWriter, r *http.Request) (err model.ErrorModel, output out.PayloadResponseSuccess) {
	var (
		fileName        = "GetListRecruitmentService.go"
		funcName        = "GetListRecruitmentService"
		inputStruct     in.RecruitmentRequest
		errFetch        error
		responseData    []byte
		urlLink         *url.URL
		responseObjects []out.RecruitmentResponse
	)

	inputStruct = input.readGetListData(r)

	defer func() {
		if errFetch != nil {
			logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
			logModel.Code = http.StatusInternalServerError
			logModel.Message = errFetch.Error()
			util.LogError(logModel.LoggerZapFieldObject())
		}
	}()

	urlLink, errFetch = url.Parse(fmt.Sprintf(`http://dev3.dansmultipro.co.id/api/recruitment/positions.json`))
	if errFetch != nil {
		err = model.GenerateErrorModel(http.StatusInternalServerError, "Kesalahan pada system, hubungi CS kami", fileName, funcName)
		return
	}

	queryUrlLink := urlLink.Query()

	if inputStruct.Page > 0 {
		queryUrlLink.Set("page", strconv.Itoa(inputStruct.Page))
	}

	if inputStruct.Description != "" {
		queryUrlLink.Set("description", inputStruct.Description)
	}

	if inputStruct.Location != "" {
		queryUrlLink.Set("location", inputStruct.Location)
	}

	if inputStruct.FullTimeSet {
		queryUrlLink.Set("full_time", strconv.FormatBool(inputStruct.FullTime))
	}

	urlLink.RawQuery = queryUrlLink.Encode()
	linkAPI := urlLink.String()

	responseData, errFetch, err = input.fetchDataHitAPI(linkAPI)
	if errFetch != nil {
		return
	}

	_ = json.Unmarshal(responseData, &responseObjects)

	output.Message = "Berhasil Get List Data Data"
	output.Data = responseObjects
	err = model.GenerateNonErrorModel()
	return
}
