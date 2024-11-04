package http

import (
	"errors"
	"github.com/04Akaps/trading_bot.git/common/json"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/go-resty/resty/v2"
	"time"
)

const (
	_paramLength = "get param length and request length mis match"
)

type client struct {
	*resty.Client
}

var HttpClient client

func init() {

	restyC := resty.New().
		SetJSONMarshaler(json.JsonHandler.Marshal).
		SetJSONUnmarshaler(json.JsonHandler.Unmarshal).
		SetTimeout(10 * time.Second)

	HttpClient = client{restyC}
}

func (u client) POST(url string, req, resp interface{}) error {
	body, err := HttpClient.JSONMarshal(req)

	if err != nil {
		return err
	}

	_, err = HttpClient.R().
		SetBody(body).
		SetResult(&resp).
		Post(url)

	if err != nil {
		return err
	}

	return nil
}

func (u client) GET(url string, paramName, req []string, resp interface{}) error {
	if len(paramName) != len(req) {
		return errors.New(_paramLength)
	}

	if len(paramName) != 0 {
		url += "?"
	}

	for i, v := range paramName {
		if i != 0 {
			url += "&"
		}
		url += v + "=" + req[i]
	}

	_, err := HttpClient.R().
		SetResult(&resp).
		Get(url)

	if err != nil {
		return err
	}

	return nil
}

func (u client) GetCurrentPriceTicker(url, headerKey, apiKey string, buffer []*types.CurrentPriceTicker) error {
	_, err := HttpClient.R().SetHeader(headerKey, apiKey).SetResult(buffer).Get(url)

	return err
}
