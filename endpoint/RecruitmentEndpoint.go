package endpoint

import (
	"github.com/alif-github/project-test/service/RecruitmentService"
	"net/http"
)

type recruitmentEndpoint struct {
	AbstractEndpoint
}

var RecruitmentEndpoint = recruitmentEndpoint{}.Initiate()

func (input recruitmentEndpoint) Initiate() (output recruitmentEndpoint) {
	input.FileName = "RecruitmentEndpoint.go"
	return
}

func (input recruitmentEndpoint) GetDetailPosition(response http.ResponseWriter, request *http.Request) {
	input.ServeJWTTokenValidationEndpoint(response, request, RecruitmentService.RecruitmentService.GetDetailRecruitmentService)
}
