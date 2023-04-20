package RecruitmentService

import (
	"github.com/alif-github/project-test/dto/in"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/service"
	"github.com/alif-github/project-test/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

type recruitmentService struct {
	service.AbstractService
}

var RecruitmentService = recruitmentService{}.Initiate()

func (input recruitmentService) Initiate() (output recruitmentService) {
	output.FileName = "RecruitmentService.go"
	return
}

func (input recruitmentService) readGetListData(request *http.Request) (inputStruct in.RecruitmentRequest) {
	inputStruct.Page, _ = strconv.Atoi(util.GenerateQueryValue(request.URL.Query()["page"]))
	inputStruct.Description = util.GenerateQueryValue(request.URL.Query()["description"])
	inputStruct.Location = util.GenerateQueryValue(request.URL.Query()["location"])
	fullTimeStr := util.GenerateQueryValue(request.URL.Query()["full_time"])
	if fullTimeStr != "" {
		inputStruct.FullTimeSet = true
		inputStruct.FullTime, _ = strconv.ParseBool(fullTimeStr)
	}

	return
}

func (input recruitmentService) fetchDataHitAPI(url string) (responseData []byte, errFetch error, err model.ErrorModel) {
	var (
		funcName = "fetchDataHitAPI"
		response *http.Response
	)

	defer func() {
		if errFetch != nil {
			err = model.GenerateErrorModel(http.StatusInternalServerError, "Kesalahan pada system, hubungi CS kami", input.FileName, funcName)
			return
		}
	}()

	response, errFetch = http.Get(url)
	if errFetch != nil {
		return
	}

	responseData, errFetch = ioutil.ReadAll(response.Body)
	if errFetch != nil {
		return
	}

	err = model.GenerateNonErrorModel()
	return
}
