package event

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"

	eventModel "github.com/pokotyan/connpass-map-api/domain/connpass/event"
	"github.com/pokotyan/connpass-map-api/infrastructure/connpass/request"
)

type Res struct {
	ResultsReturned  int                `json:results_returned`
	ResultsStart     int                `json:results_start`
	ResultsAvailable int                `json:results_available`
	Events           []eventModel.Event `json:events`
}

type ReqParam struct {
	Keyword string
	YmList  []int
	YmdList []int
}

func addUrlVal(urlVal url.Values, value []int) func(string) {
	return func(conppassParam string) {
		for _, v := range value {
			urlVal.Add(conppassParam, strconv.Itoa(v))
		}
	}
}

type Param map[string]interface{}

func createBuildParam(param Param) func() string {
	values := url.Values{}

	return func() string {
		for key, value := range param {
			switch key {
			case "ymdList":
				addUrlVal(values, value.([]int))("ymd")
			case "ymList":
				addUrlVal(values, value.([]int))("ym")
			case "keyword":
				if value.(string) != "" { // keyword指定していなければ（ゼロ値だったら）addしない
					values.Add(key, value.(string))
				}
			}
		}

		return values.Encode()
	}
}

func Get(reqParam ReqParam) Res {
	param := map[string]interface{}{
		"keyword": reqParam.Keyword,
		"ymList":  reqParam.YmList,
		"ymdList": reqParam.YmdList,
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
