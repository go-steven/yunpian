package yunpian

type SendResult struct {
	Response *SmsSendResponse
	Err      error
}

type SmsMessage struct {
	mobiles []string         `json:"mobiles,omitempty"` // mobile列表
	text    string           `json:"text,omitempty"`
	retChan chan *SendResult `json:"-"`
}

func NewSmsMessage(mobiles []string, text string) *SmsMessage {
	return &SmsMessage{
		mobiles: mobiles,
		text:    text,
		retChan: make(chan *SendResult, 1),
	}
}

func (m *SmsMessage) Result() *SendResult {
	return <-m.retChan
}

type SmsService struct {
	apiKey      string
	callbackUrl string

	msgs     chan *SmsMessage
	exitChan chan struct{}
}

func NewSmsService(smsApiKey, callbackUrl string, queueLen uint) *SmsService {
	if queueLen < 100 {
		queueLen = 100
	}

	return &SmsService{
		apiKey:      smsApiKey,
		callbackUrl: callbackUrl,

		msgs:     make(chan *SmsMessage, queueLen),
		exitChan: make(chan struct{}, 1),
	}
}

func (s *SmsService) Send(m *SmsMessage) *SendResult {
	s.msgs <- m
	return m.Result()
}

func (s *SmsService) Start() {
	for {
		select {
		case m := <-s.msgs:
			m.retChan <- s.notify(m)
		case <-s.exitChan:
			return
		}
	}
}

func (s *SmsService) Stop() {
	s.exitChan <- struct{}{}
}

func (s *SmsService) notify(m *SmsMessage) *SendResult {
	apiReq := &SmsSendRequest{}
	apiReq.ApiKey = s.apiKey
	apiReq.Mobiles = m.mobiles
	apiReq.Text = m.text
	if s.callbackUrl != "" {
		apiReq.CallbackUrl = s.callbackUrl
	}
	ret, err := SendSms(apiReq)
	return &SendResult{
		Response: ret,
		Err:      err,
	}
}

func (s *SmsService) AddSign(sign string) (map[string]interface{}, error) {
	apiReq := &AddSignRequest{}
	apiReq.ApiKey = s.apiKey
	apiReq.Sign = sign
	ret, err := AddSign(apiReq)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SmsService) GetSigns(sign string) ([]*Sign, error) {
	apiReq := &GetSignRequest{}
	apiReq.ApiKey = s.apiKey
	ret, err := GetSign(apiReq)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SmsService) GetUser(sign string) (*GetUserResponse, error) {
	apiReq := &GetUserRequest{}
	apiReq.ApiKey = s.apiKey
	ret, err := GetUser(apiReq)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
