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

////Transaction Status

type TransactionStatusRequestBody struct {
	Initiator          string
	SecurityCredential string
	ShortCode          string

	CommandID                string
	TransactionID            string
	OriginatorConversationID string
	PartyA                   string
	IdentifierType           string
	Remarks                  string
	ResultURL                string
	QueueTimeOutURL          string
	Occasion                 string
}

/////Register C2B url

type RegisterC2BURLRequestBody struct {
	ValidationURL   string
	ConfirmationURL string
	ResponseType    string
	ShortCode       string
}

type C2BSimulationRequestBody struct {
	CommandID string
	Amount    string
	///phone number
	Msisdn string
	///optional
	BillRefNumber string
	ShortCode     string
}

type C2BValidationResponse struct {
	BillRefNumber     string
	BusinessShortCode string
	FirstName         string
	InvoiceNumber     string
	LastName          string
	MSISDN            string
	MiddleName        string
	OrgAccountBalance string
	ThirdPartyTransID string
	TransAmount       string
	TransID           string
	TransTime         string
	TransactionType  string
}

type MpesaResult struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

type MpesaTransactionStatus struct {
	//////
	Result MpesaTransactionResult
}

type MpesaTransactionResult struct {
	Result
	ReferenceData BalanceReferenceData
}

type MpesaBalance struct {
	Result BalanceResult `json:"Result"`
}

type BalanceReferenceData struct {
	ReferenceItem ReferenceItem `json:"ReferenceItem"`
}
type BalanceResult struct {
	Result

	ReferenceData BalanceReferenceData `json:"ReferenceData"`
}

type MpesaB2BResponse struct {
	Result MpesaB2BResult `json:"Result"`
}

type B2BReferenceData struct {
	ReferenceItem []ReferenceItem `json:"ReferenceItem"`
}
type MpesaB2BResult struct {
	Result

	ReferenceData B2BReferenceData `json:"ReferenceData"`
}

type MpesaResponse struct {
	Result Result `json:"Result"`
}
type ReferenceItem struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value,omitempty"`
}
type ReferenceData struct {
	ReferenceItem []ReferenceItem `json:"ReferenceItem"`
}
type ResultParameter struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value,omitempty"`
}
type ResultParameters struct {
	ResultParameter []ResultParameter `json:"ResultParameter"`
}
type Result struct {
	ConversationID string `json:"ConversationID"`
	ResultType     int    `json:"ResultType"`

	ResultCode int    `json:"ResultCode"`
	ResultDesc string `json:"ResultDesc"`

	OriginatorConversationID string `json:"OriginatorConversationID"`

	TransactionID string `json:"TransactionID"`

	ResultParameters ResultParameters `json:"ResultParameters"`

	//ReferenceData            ReferenceData    `json:"ReferenceData"`

}
