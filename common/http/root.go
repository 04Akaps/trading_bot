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

func NewClient(headerKey, apiKey string) client {
	restyC := resty.New().
		SetJSONMarshaler(json.JsonHandler.Marshal).
		SetJSONUnmarshaler(json.JsonHandler.Unmarshal).
		SetTimeout(10*time.Second).
		SetHeader(headerKey, apiKey)

	return client{restyC}
}

func (u client) POST(url string, req, resp interface{}) error {
	body, err := u.JSONMarshal(req)

	if err != nil {
		return err
	}

	_, err = u.R().
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

	_, err := u.R().
		SetResult(&resp).
		Get(url)

	if err != nil {
		return err
	}

	return nil
}

func (u client) GetCurrentPriceTicker(url string, buffer *[]*types.CurrentPriceTicker) error {
	_, err := u.R().SetResult(buffer).Get(url)
	return err
}

func (u client) GetTradingDay(url string, buffer *[]*types.VolumeTicker) error {
	_, err := u.R().SetResult(buffer).Get(url)
	return err
}
