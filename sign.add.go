package yunpian

import (
	"github.com/bububa/ljson"
)

type AddSignRequest struct {
	BaseRequest

	Sign   string
	Notify bool
}

func AddSign(apiReq *AddSignRequest) (map[string]interface{}, error) {
	client := NewClient(apiReq.ApiKey)
	req := NewRequest("sign/add.json")
	req.Params["sign"] = apiReq.Sign
	req.Params["notify"] = apiReq.Notify

	response, err := client.Execute(req)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{})
	err = ljson.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
