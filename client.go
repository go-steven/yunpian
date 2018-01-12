package yunpian

import (
	"errors"
	"fmt"
	"github.com/bububa/ljson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	GATEWAY_URL = "https://sms.yunpian.com/v2/"
)

type Request struct {
	MethodUrl string
	Params    map[string]interface{}
}

func NewRequest(methodUrl string) *Request {
	return &Request{
		MethodUrl: methodUrl,
		Params:    make(map[string]interface{}),
	}
}

type Client struct {
	ApiKey string
}

//create new client
func NewClient(apiKey string) (c *Client) {
	c = &Client{
		ApiKey: apiKey,
	}
	return
}

func (c *Client) Execute(req *Request) ([]byte, error) {
	sysParams := make(map[string]string)
	sysParams["apikey"] = c.ApiKey
	for k, v := range req.Params {
		switch v.(type) {
		case string:
			sysParams[k] = v.(string)
		case uint64:
			sysParams[k] = fmt.Sprintf("%d", v.(uint64))
		case uint32:
			sysParams[k] = fmt.Sprintf("%d", v.(uint32))
		case uint16:
			sysParams[k] = fmt.Sprintf("%d", v.(uint16))
		case uint8:
			sysParams[k] = fmt.Sprintf("%d", v.(uint8))
		case uint:
			sysParams[k] = fmt.Sprintf("%d", v.(uint))
		case int64:
			sysParams[k] = fmt.Sprintf("%d", v.(int64))
		case int32:
			sysParams[k] = fmt.Sprintf("%d", v.(int32))
		case int16:
			sysParams[k] = fmt.Sprintf("%d", v.(int16))
		case int8:
			sysParams[k] = fmt.Sprintf("%d", v.(int8))
		case int:
			sysParams[k] = fmt.Sprintf("%d", v.(int))
		case float32:
			sysParams[k] = fmt.Sprintf("%f", v.(float32))
		case float64:
			sysParams[k] = fmt.Sprintf("%f", v.(float64))
		case bool:
			if v.(bool) {
				sysParams[k] = "true"
			} else {
				sysParams[k] = "false"
			}
		}
	}
	values := url.Values{}
	for k, v := range sysParams {
		values.Add(k, v)
	}
	fmt.Println(Json(sysParams))
	reqUrl := GATEWAY_URL + req.MethodUrl
	response, err := http.DefaultClient.Post(reqUrl, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	j := make(map[string]interface{})
	err = ljson.Unmarshal(body, &j)
	if err == nil {
		code, err := strconv.ParseInt(fmt.Sprintf("%v", j["code"]), 10, 64)
		if err == nil && code != 0 {
			return nil, errors.New(string(body))
		}
	}
	fmt.Println(Json(j))
	return body, nil
}
