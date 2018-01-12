package yunpian

import (
	"fmt"
	"github.com/bububa/ljson"
)

type GetSignRequest struct {
	BaseRequest

	Sign string
}

type Sign struct {
	Chan         string `json:"chan" codec:"chan"`
	CheckStatus  string `json:"check_status" codec:"check_status"`
	Enabled      bool   `json:"enabled" codec:"enabled"`
	Extend       string `json:"extend" codec:"extend"`
	IndustryType string `json:"industry_type" codec:"industry_type"`
	OnlyGlobal   bool   `json:"only_global" codec:"only_global"`
	Remark       string `json:"remark" codec:"remark"`
	Sign         string `json:"sign" codec:"sign"`
	Vip          bool   `json:"vip" codec:"vip"`
}

type GetSignResponse struct {
	Code  int     `json:"code" codec:"code"`
	Total int     `json:"msg" codec:"total"`
	Sign  []*Sign `json:"sign" codec:"sign"`
}

func GetSign(apiReq *GetSignRequest) ([]*Sign, error) {
	client := NewClient(apiReq.ApiKey)
	req := NewRequest("sign/get.json")
	if apiReq.Sign != "" {
		req.Params["sign"] = fmt.Sprintf("【%s】", apiReq.Sign)
	}

	response, err := client.Execute(req)
	if err != nil {
		return nil, err
	}
	ret := GetSignResponse{}
	err = ljson.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Sign, nil
}
