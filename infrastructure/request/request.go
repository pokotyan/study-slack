package request

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func New(baseURL string) (*Client, error) {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}, nil
}

func (c *Client) Get(
	path string,
	buildParam func() string,
) ([]byte, error) {
	url := c.BaseURL + path

	req, _ := http.NewRequest("GET", url, nil)
	req.URL.RawQuery = buildParam()

	return c.request(req)
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
