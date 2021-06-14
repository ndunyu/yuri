package yuri

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}


type StkPushResult struct {
	///
	CheckoutRequestID string
	CustomerMessage string
	MerchantRequestID string
	ResponseCode string
	ResponseDescription string

}






type StkPushResponseBody struct {
	Body Body `json:"Body"`
}
type Item struct {
	Name  string `json:"Name"`
	Value interface{}   `json:"Value,omitempty"`
}
type CallbackMetadata struct {
	Item []Item `json:"Item"`
}
type StkCallback struct {
	MerchantRequestID string           `json:"MerchantRequestID"`
	CheckoutRequestID string           `json:"CheckoutRequestID"`
	ResultCode        int              `json:"ResultCode"`
	ResultDesc        string           `json:"ResultDesc"`
	CallbackMetadata  CallbackMetadata `json:"CallbackMetadata"`
}
type Body struct {
	StkCallback StkCallback `json:"stkCallback"`
}
///when querying for success/failure
type StkPushQueryRequestBody struct {
	BusinessShortCode	string
	Password	string
	Timestamp	string
	CheckoutRequestID	string
}

type StkPushQueryResponseBody struct {
	MerchantRequestID	string
	CheckoutRequestID	string
	ResponseCode	string
	ResultDesc	string
	ResponseDescription	string
	ResultCode	string

}





type StKPushRequestBody struct {
	BusinessShortCode string
	Password string
	Timestamp string
	///use only [ CustomerPayBillOnline ]
	TransactionType string
	Amount string
	//sender phone number
	PartyA string
	///receiver shortcode
	PartyB string
	////Sending funds
	PhoneNumber string
	///
	CallBackURL string
	///use this with paybill
	AccountReference string
	//
	TransactionDesc string


}
// MpesaResponse is returned by every mpesa api
// Here
// i.e that is when we call Mpesa.sendAndProcessMpesaRequest
type MpesaResult struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorCoversationID string `json:"OriginatorCoversationID"`
	/// OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription   string `json:"ResponseDescription"`
	///ResponseDescription      string `json:"ResponseDescription"`
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
	Initiator string
	//see https://developer.safaricom.co.ke/docs#security-credentials
	SecurityCredential string
	// You don't have to pass
	//this we have already set it to AccountBalance
	CommandID string
	// this is your shortcode i.e Paybill or till number
	PartyA string
	// use PayBillIdentifier for paybill
	// use TillNumberIdentifier for TillNumber
	IdentifierType string
	Remarks        string
	// url to receive the results
	ResultURL       string
	QueueTimeOutURL string
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

type C2BValidationAndConfirmationResponse struct {
	BillRefNumber     string
	BusinessShortCode string
	FirstName         string
	InvoiceNumber     string
	LastName          string
	MSISDN            string
	MiddleName        string
	OrgAccountBalance string
	ThirdPartyTransID string
	TransAmount        float64 `json:",string"`
	TransID           string
	TransTime         string
	TransactionType   string
}

/// this is what
// C2BValidatedResponse you send back to mpesa after receiving the validation
//from c2B api
//
// request to Time confirm if it is okay
//
type C2BValidatedResponse struct {
	// use 0 for success
	// and any other value for error
	// this is required
	ResultCode int
	// this is optional
	// value can be  “Service processing successful”
	ResultDesc string
	///The unique identifier of the payment transaction that is generated by the third party
	ThirdPartyTransID string
}

// This is response sent when we are checkig the status of a transaction
type MpesaTransactionStatus struct {
	//////
	Result MpesaTransactionResult
}

type MpesaTransactionResult struct {
	Result
	ReferenceData BalanceReferenceData
}

// MpesaBalance is what is returned to the resulturl
// you set the result url in your AccountBalanceRequestBody.ResultURL
// you query for your balance using the balance
// api  ....
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

//end of balance result

// MpesaB2BResponse is what is returned to the B2BRequestBody.ResultURL
// you set the result url in your AccountBalanceRequestBody.ResultURL
// This is when using the B2B api
type MpesaB2BResponse struct {
	Result MpesaB2BResult `json:"Result"`
}
type MpesaB2BResult struct {
	Result

	ReferenceData B2BReferenceData `json:"ReferenceData"`
}
type B2BReferenceData struct {
	ReferenceItem []ReferenceItem `json:"ReferenceItem"`
}

////End of B2B result

type MpesaResponse struct {
	Result Result `json:"Result"`
}
type Result struct {
	ConversationID string `json:"ConversationID"`
	ResultType     int    `json:"ResultType"`

	ResultCode string    `json:"ResultCode"`
	ResultDesc string `json:"ResultDesc"`

	OriginatorConversationID string `json:"OriginatorConversationID"`

	TransactionID string `json:"TransactionID"`

	ResultParameters ResultParameters `json:"ResultParameters"`

	//ReferenceData            ReferenceData    `json:"ReferenceData"`

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
