package yunpian

import (
	"github.com/bububa/ljson"
)

type UpdateTplRequest struct {
	BaseRequest
	TplId      uint64
	TplContent string
}

func UpdateTpl(apiReq *UpdateTplRequest) (*GetTplResponse, error) {
	client := NewClient(apiReq.ApiKey)
	req := NewRequest("tpl/update.json")
	req.Params["tpl_id"] = apiReq.TplId
	req.Params["tpl_content"] = apiReq.TplContent

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
