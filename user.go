package yunpian

import (
	"github.com/bububa/ljson"
)

type GetUserRequest struct {
	BaseRequest
}

type GetUserResponse struct {
	Nick             string      `json:"nick" codec:"nick"`
	GmtCreated       string      `json:"gmt_created" codec:"gmt_created"`
	Mobile           string      `json:"mobile" codec:"mobile"`
	IpWhitelist      interface{} `json:"ip_whitelist" codec:"ip_whitelist"`
	ApiVersion       string      `json:"api_version" codec:"api_version"`
	Balance          uint64      `json:"balance" codec:"balance"`
	AlarmBalance     uint64      `json:"alarm_balance" codec:"alarm_balance"`
	EmergencyContact string      `json:"emergency_contact" codec:"emergency_contact"`
	EmergencyMobile  string      `json:"emergency_mobile" codec:"emergency_mobile"`
}

func GetUser(apiReq *GetUserRequest) (*GetUserResponse, error) {
	client := NewClient(apiReq.ApiKey)
	req := NewRequest("user/get.json")
	response, err := client.Execute(req)
	if err != nil {
		return nil, err
	}
	ret := GetUserResponse{}
	err = ljson.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
