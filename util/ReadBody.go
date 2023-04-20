package util

import (
	"github.com/alif-github/project-test/model"
	"io/ioutil"
	"net/http"
)

func ReadBody(request *http.Request) (byteBody []byte, bodySize int, err model.ErrorModel) {
	var (
		fileName = "ReadBody.go"
		funcName = "ReadBody"
		errS     error
	)

	defer func() {
		_ = request.Body.Close()
	}()

	byteBody, errS = ioutil.ReadAll(request.Body)
	if errS != nil {
		err = model.GenerateErrorModel(500, errS.Error(), fileName, funcName)
		return
	}

	return
}
