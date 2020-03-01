package event

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"

	"github.com/pokotyan/study-slack/infrastructure/connpass/request"
)

type (
	connpassEvent struct {
		Keyword string
		YmList  []int
		YmdList []int
	}

	Param map[string]interface{}
)

func NewConnpassEvent() ConnpassEvent {
	return &connpassEvent{}
}

func addURLVal(urlVal url.Values, value []int) func(string) {
	return func(conppassParam string) {
		for _, v := range value {
			urlVal.Add(conppassParam, strconv.Itoa(v))
		}
	}
}

func createBuildParam(param Param) func() string {
	values := url.Values{}

	return func() string {
		for key, value := range param {
			switch key {
			case "ymdList":
				addURLVal(values, value.([]int))("ymd")
			case "ymList":
				addURLVal(values, value.([]int))("ym")
			case "keyword":
				if value.(string) != "" { // keyword指定していなければ（ゼロ値だったら）addしない
					values.Add(key, value.(string))
				}
			case "count":
				values.Add(key, value.(string))
			}
		}

		return values.Encode()
	}
}

func (ce connpassEvent) Get(reqParam ReqParam) Res {
	param := map[string]interface{}{
		"keyword": reqParam.Keyword,
		"ymList":  reqParam.YmList,
		"ymdList": reqParam.YmdList,
		"count":   "100",
	}
	buildParam := createBuildParam(param)

	res, err := request.Get("/event/", buildParam)
	if err != nil {
		log.Fatal(err)
	}

	var resModel Res
	json.Unmarshal(res, &resModel)

	return resModel
}
