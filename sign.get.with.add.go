package yunpian

import (
	"errors"
)

func GetSignWithAdd(apiReq *GetSignRequest) (*Sign, error) {
	if apiReq.Sign == "" {
		return nil, errors.New("empty sign.")
	}
	signs, err := GetSign(apiReq)
	if err != nil {
		return nil, err
	}
	if len(signs) > 0 {
		return signs[0], nil
	} else {
		req := &AddSignRequest{}
		req.ApiKey = apiReq.ApiKey
		req.Sign = apiReq.Sign
		_, err := AddSign(req)
		if err != nil {
			return nil, err
		}

		signs, err := GetSign(apiReq)
		if err != nil {
			return nil, err
		}
		if len(signs) == 0 {
			return nil, errors.New("no sign.")
		}
		return signs[0], nil
	}
}
