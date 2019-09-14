package request

import (
	"log"

	Request "github.com/pokotyan/connpass-map-api/infrastructure/request"
)

type Params map[string]interface{}

func Get(
	path string,
	buildParamfun func(map[string]interface{}) string,
	param Params) ([]byte, error,
) {
	request, err := Request.New("https://connpass.com/api/v1")
	if err != nil {
		log.Fatal(err)
	}

	res, err := request.Get(path, buildParamfun)(param)
	if err != nil {
		log.Fatal(err)
	}

	return res, err
}
