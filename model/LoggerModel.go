package model

import (
	"go.uber.org/zap"
)

type LoggerModel struct {
	RequestID   string `json:"requestID"`
	Class       string `json:"class"`
	Application string `json:"application"`
	Version     string `json:"version"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
}

func GenerateLogModel(version string, application string) (output LoggerModel) {
	output.RequestID = "-"
	output.Class = "-"
	output.Application = application
	output.Version = version
	output.Code = 0
	output.Message = "-"
	return
}

func (object LoggerModel) LoggerZapFieldObject() (output []zap.Field) {
	output = append(output, zap.String("requestID", object.RequestID))
	output = append(output, zap.String("class", object.Class))
	output = append(output, zap.String("application", object.Application))
	output = append(output, zap.String("version", object.Version))
	output = append(output, zap.Int("code", object.Code))
	output = append(output, zap.String("message", object.Message))
	return
}
