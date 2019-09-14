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

func buildParam(param map[string]interface{}) string {
	values := url.Values{}

	for key, value := range param {
		switch key {
		case "ymdList":
			addUrlVal(values, value.([]int))("ymd")
		case "ymList":
			addUrlVal(values, value.([]int))("ym")
		case "keyword":
			if value.(string) != "" { // ゼロ値だったら（指定していなければ）addしない
				values.Add(key, value.(string))
			}
		}
	}

	return values.Encode()
}

func Get(reqParam ReqParam) Res {
	param := map[string]interface{}{
		"keyword": reqParam.Keyword,
		"ymList":  reqParam.YmList,
		"ymdList": reqParam.YmdList,
	}

	res, err := request.Get("/event/", buildParam, param)
	if err != nil {
		log.Fatal(err)
	}

	var resModel Res
	json.Unmarshal(res, &resModel)

	return resModel
}
