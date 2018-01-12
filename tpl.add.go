package yunpian

import (
	"github.com/bububa/ljson"
)

type AddTplRequest struct {
	BaseRequest
	TplContent string
}

func AddTpl(apiReq *AddTplRequest) (*GetTplResponse, error) {
	client := NewClient(apiReq.ApiKey)
	req := NewRequest("tpl/add.json")
	req.Params["tpl_content"] = apiReq.TplContent
	req.Params["notify_type"] = 3

	response, err := client.Execute(req)
	if err != nil {
		return nil, err
	}
	ret := GetTplResponse{}
	err = ljson.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
