package RecruitmentService

import "github.com/alif-github/project-test/service"

type recruitmentService struct {
	service.AbstractService
}

var RecruitmentService = recruitmentService{}.Initiate()

func (input recruitmentService) Initiate() (output recruitmentService) {
	output.FileName = "RecruitmentService.go"
	return
}
