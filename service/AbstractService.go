package service

import (
	"encoding/json"
	"github.com/alif-github/project-test/constanta"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/util"
	"net/http"
)

type AbstractService struct {
	FileName string
}

func (input AbstractService) ReadInput(r *http.Request, inputStruct interface{}) (bodySize int, err model.ErrorModel) {
	var (
		fileName = "AbstractService.go"
		funcName = "ReadInput"
		errS     error
		byteReq  []byte
	)

	if r.Method != constanta.GetStatus {
		byteReq, bodySize, err = util.ReadBody(r)
		if err.Error != nil {
			return
		}
	}

	if string(byteReq) != "" {
		errS = json.Unmarshal(byteReq, &inputStruct)
		if errS != nil {
			err = model.GenerateErrorModel(500, errS.Error(), fileName, funcName)
			return
		}
	}

	return
}
