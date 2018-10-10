package formatter

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

//Jsend used to format JSON with jsend rules
type Jsend struct {
	Status  string      `json:"status" binding:"required"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func FailResponse(msg string) Jsend {
	return Jsend{Status: "failed", Message: msg}
}

func SuccessResponse() Jsend {
	return Jsend{Status: "success"}
}

func ObjectResponse(data interface{}) Jsend {
	return Jsend{Status: "success", Data: data}
}

func (j Jsend) Stringify() string {
	toJSON, err := json.Marshal(j)
	if err != nil {
		logrus.WithError(err).Errorln("Unable to stringify JSON")
		return ""
	}
	return string(toJSON)
}
