package yunpian

import (
	"github.com/bububa/ljson"
	"strings"
)

type SmsSendRequest struct {
	BaseRequest

	Mobiles     []string
	Text        string
	CallbackUrl string
}

type SmsSendResponse struct {
	TotalCount uint64           `json:"total_count" codec:"total_count"`
	TotalFee   string           `json:"total_fee" codec:"total_fee"`
	Unit       string           `json:"unit" codec:"unit"`
	Data       []*SmsSendResult `json:"data" codec:"data"`
}
type SmsSendResult struct {
	Code   uint8   `json:"code" codec:"code"`
	Msg    string  `json:"msg" codec:"msg"`
	Count  uint64  `json:"count" codec:"count"`
	Fee    float64 `json:"fee" codec:"fee"`
	Unit   string  `json:"unit" codec:"unit"`
	Mobile string  `json:"mobile" codec:"mobile"`
	Sid    uint64  `json:"sid" codec:"sid"`
}

func SendSms(apiReq *SmsSendRequest) (*SmsSendResponse, error) {
	ret := SmsSendResponse{
		TotalCount: 0,
		TotalFee:   "0.0",
		Unit:       "RMB",
		Data:       []*SmsSendResult{},
	}

	if len(apiReq.Mobiles) == 0 {
		return &ret, nil
	}

	client := NewClient(apiReq.ApiKey)
	req := NewRequest("sms/batch_send.json")
	req.Params["mobile"] = strings.Join(apiReq.Mobiles, ",")
	req.Params["text"] = apiReq.Text
	if apiReq.CallbackUrl != "" {
		req.Params["callback_url"] = apiReq.CallbackUrl
	}
	response, err := client.Execute(req)
	if err != nil {
		return nil, err
	}

	err = ljson.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
