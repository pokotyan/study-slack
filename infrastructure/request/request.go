package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func New(baseUrl string) (*Client, error) {
	return &Client{
		BaseURL:    baseUrl,
		HTTPClient: &http.Client{},
	}, nil
}

func (c *Client) Get(
	path string,
	buildParamfun func (map[string]interface{}) string,
) func(param map[string]interface{}) ([]byte, error) {
	url := c.BaseURL + path

	req, _ := http.NewRequest("GET", url, nil)

	return func(param map[string]interface{}) ([]byte, error)  {
		req.URL.RawQuery = buildParamfun(param)

		fmt.Println(req.URL.RawQuery)

		return c.request(req)
	}
}

func (c *Client) request(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
