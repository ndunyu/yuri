package yuri

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type B2CRequestBody struct {
	InitiatorName      string
	SecurityCredential string
	CommandID          string
	Amount             int
	PartyA             string
	PartyB             string
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
	Occassion          string
}

type B2BRequestBody struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	Amount                 int
	PartyA                 string
	SenderIdentifierType   string
	RecieverIdentifierType string
	AccountReference       string
	PartyB                 string
	Remarks                string
	QueueTimeOutURL        string
	ResultURL              string
}

type AccountBalanceRequestBody struct {
	Initiator          string
	SecurityCredential string
	CommandID          string
	PartyA             string
	IdentifierType     string
	Remarks            string
	ResultURL          string
	QueueTimeOutURL    string
}

type MpesaResult struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

type MpesaResponse struct {
	Result Result `json:"Result"`
}
type ReferenceItem struct {
	Key   string `json:"Key"`
	Value int64  `json:"Value,omitempty"`
}
type ReferenceData struct {
	ReferenceItem []ReferenceItem `json:"ReferenceItem"`
}
type ResultParameter struct {
	Key   string `json:"Key"`
	Value interface{} `json:"Value,omitempty"`
}
type ResultParameters struct {
	ResultParameter []ResultParameter `json:"ResultParameter"`
}
type Result struct {
	ConversationID           string           `json:"ConversationID"`
	OriginatorConversationID string           `json:"OriginatorConversationID"`
	ReferenceData            ReferenceData    `json:"ReferenceData"`
	ResultCode               int              `json:"ResultCode"`
	ResultDesc               string           `json:"ResultDesc"`
	ResultParameters         ResultParameters `json:"ResultParameters"`
	ResultType               int              `json:"ResultType"`
	TransactionID            string           `json:"TransactionID"`



}



