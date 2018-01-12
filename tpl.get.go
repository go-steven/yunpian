package yunpian

import (
	"github.com/bububa/ljson"
)

type GetTplRequest struct {
	BaseRequest
}

type GetTplResponse struct {
	TplId       uint64      `json:"tpl_id" codec:"tpl_id"`
	TplContent  string      `json:"tpl_content" codec:"tpl_content"`
	CheckStatus string      `json:"check_status" codec:"check_status"`
	Reason      interface{} `json:"reason" codec:"reason"`
}

func GetTpl(apiReq *GetTplRequest) ([]*GetTplResponse, error) {
	client := NewClient(apiReq.ApiKey)
	req := NewRequest("tpl/get.json")
	response, err := client.Execute(req)
	if err != nil {
		return nil, err
	}
	ret := []*GetTplResponse{}
	err = ljson.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
