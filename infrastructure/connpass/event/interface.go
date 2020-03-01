package event

import (
	eventModel "github.com/pokotyan/study-slack/domain/connpass/model/event"
)

type (
	ReqParam struct {
		Keyword string
		YmList  []int
		YmdList []int
	}

	Res struct {
		ResultsReturned  int                `json:"results_returned"`
		ResultsStart     int                `json:"results_start"`
		ResultsAvailable int                `json:"results_available"`
		Events           []eventModel.Event `json:"events"`
	}

	ConnpassEvent interface {
		Get(reqParam ReqParam) Res
	}
)
