package model

import "errors"

type ErrorModel struct {
	Code     int
	Error    error
	FileName string
	FuncName string
}

func GenerateErrorModel(code int, err string, fileName string, funcName string) (errModel ErrorModel) {
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.FileName = fileName
	errModel.FuncName = funcName
	return
}

func GenerateNonErrorModel() (errModel ErrorModel) {
	errModel.Code = 200
	errModel.Error = nil
	return
}
