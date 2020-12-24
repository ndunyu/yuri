package yuri






type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn  string `json:"expires_in"`

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
	Initiator      string
	SecurityCredential string
	CommandID          string
	Amount             int
	PartyA             string
	SenderIdentifierType  string
	RecieverIdentifierType string
	AccountReference string
	PartyB             string
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
}




type Result struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}



