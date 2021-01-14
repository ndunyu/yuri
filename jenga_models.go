package yuri

type JengaAccessToken struct {
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	ExpiresIn   string `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

///Start of indentity
type IdentityRequestBody struct {
	Identity Identity `json:"identity"`
}
type Identity struct {
	DocumentType   string `json:"documentType"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	DateOfBirth    string `json:"dateOfBirth"`
	DocumentNumber string `json:"documentNumber"`
	CountryCode    string `json:"countryCode"`
}

type IdentityResponseBody struct {
	Identity KycResponseIdentity `json:"identity"`
}
type Customer struct {
	FullName      string `json:"fullName"`
	FirstName     string `json:"firstName"`
	Middlename    string `json:"middlename"`
	LastName      string `json:"lastName"`
	ShortName     string `json:"ShortName"`
	BirthDate     string `json:"birthDate"`
	BirthCityName string `json:"birthCityName"`
	DeathDate     string `json:"deathDate"`
	Gender        string `json:"gender"`
	FaceImage     string `json:"faceImage"`
	Occupation    string `json:"occupation"`
	Nationality   string `json:"nationality"`
}
type AdditionalIdentityDetails struct {
	DocumentNumber string `json:"documentNumber"`
	DocumentType   string `json:"documentType"`
	IssuedBy       string `json:"issuedBy"`
}
type Address struct {
	ProvinceName    string `json:"provinceName"`
	DistrictName    string `json:"districtName"`
	LocationName    string `json:"locationName"`
	SubLocationName string `json:"subLocationName"`
	VillageName     string `json:"villageName"`
}
type KycResponseIdentity struct {
	Customer                  Customer                    `json:"customer"`
	DocumentType              string                      `json:"documentType"`
	DocumentNumber            string                      `json:"documentNumber"`
	DocumentSerialNumber      string                      `json:"documentSerialNumber"`
	DocumentIssueDate         string                      `json:"documentIssueDate"`
	DocumentExpirationDate    string                      `json:"documentExpirationDate"`
	IssuedBy                  string                      `json:"IssuedBy"`
	AdditionalIdentityDetails []AdditionalIdentityDetails `json:"additionalIdentityDetails"`
	Address                   Address                     `json:"address"`
}

///end of kyc identity

/// start of airtime
type AirtimeRequest struct {
	Customer AirTimeRequestCustomer `json:"customer"`
	Airtime  Airtime                `json:"airtime"`
}
type AirTimeRequestCustomer struct {
	CountryCode  string `json:"countryCode"`
	MobileNumber string `json:"mobileNumber"`
}
type Airtime struct {
	Amount    string `json:"amount"`
	Reference string `json:"reference"`
	Telco     string `json:"telco"`
}

type AirtimeResponse struct {
	ReferenceNumber string `json:"referenceNumber"`
	Status          string `json:"status"`
	ResponseStatus  string `json:"response_status"`
	ResponseMsg string `json:"response_msg"`
	ResponseCode string `json:"response_code"`

}

////end of aitime purcahse

///money transfer from bank account to mobile money

type BankToMobileMoneyRequest struct {
	Source      Source      `json:"source"`
	Destination Destination `json:"destination"`
	Transfer    Transfer    `json:"transfer"`
}
type Source struct {
	CountryCode   string `json:"countryCode"`
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
}
type Destination struct {
	Type         string `json:"type"`
	CountryCode  string `json:"countryCode"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobileNumber"`
	WalletName   string `json:"walletName"`
}
type Transfer struct {
	Type         string `json:"type"`
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
	Reference    string `json:"reference"`
	Date         string `json:"date"`
	Description  string `json:"description"`
}


////


type SendMoneyResponse struct {
	TransactionID string `json:"transactionId"`
	Status        string `json:"status"`
}
