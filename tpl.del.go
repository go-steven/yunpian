package yunpian

import (
	"github.com/bububa/ljson"
)

type DeleteTplRequest struct {
	BaseRequest
	TplId uint64
}

func DeleteTpl(apiReq *DeleteTplRequest) (*GetTplResponse, error) {
	client := NewClient(apiReq.ApiKey)
	req := NewRequest("tpl/del.json")
	req.Params["tpl_id"] = apiReq.TplId

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
