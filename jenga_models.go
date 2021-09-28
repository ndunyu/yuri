package yuri

///Jenga Balance
type JengaBalance struct {
	Currency string     `json:"currency"`
	Balances []Balances `json:"balances"`
}
type Balances struct {
	Amount float64 `json:"amount,string"`
	Type   string  `json:"type"`
}

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
	ResponseMsg     string `json:"response_msg"`
	ResponseCode    string `json:"response_code"`
}

////end of aitime purcahse

///money transfer from bank account to mobile money

type BankToMobileMoneyRequest struct {
	Source      Source                 `json:"source"`
	Destination MobileMoneyDestination `json:"destination"`
	Transfer    Transfer               `json:"transfer"`
}
type Source struct {
	CountryCode   string `json:"countryCode" validate:"required"`
	Name          string `json:"name" validate:"required"`
	AccountNumber string `json:"accountNumber" validate:"required"`
}
type Destination struct {
	Type         string `json:"type" `
	CountryCode  string `json:"countryCode" validate:"required"`
	Name         string `json:"name" validate:"required"`
	MobileNumber string `json:"mobileNumber" validate:"required"`
}
type MobileMoneyDestination struct {
	Destination
	WalletName string `json:"walletName"`
}
type Transfer struct {
	Type         string `json:"type" validate:"required"`
	Amount       string `json:"amount" validate:"required"`
	CurrencyCode string `json:"currencyCode" validate:"required"`
	Reference    string `json:"reference" validate:"required"`
	Date         string `json:"date" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

////Pesalink starts here

type PesaLinkRequest struct {
	Source      Source              `json:"source"`
	Transfer    Transfer            `json:"transfer"`
	Destination PesaLinkDestination `json:"destination"`
}
type PesaLinkDestination struct {
	Destination
	BankCode      string `json:"bankCode" validate:"required"`
	AccountNumber string `json:"accountNumber" validate:"required"`
}

///TODO::Use the validate package
///TO reduce the code
func (p *PesaLinkRequest) Validate() {
	///err = s.Validate.Struct(p)
	//Start of source validation
	///if IsEmpty(p.Source.AccountNumber) {

	//}
	//if IsEmpty(p.Source.Name) {

	//}
	//if IsEmpty(p.Source.CountryCode) {

	//}
	///End of source validation
	//if IsEmpty(p.Destination.CountryCode) {

	//}
	//if IsEmpty(p.Destination.Name) {

	//}
	//if IsEmpty(p.Destination.CountryCode) {

	//}

}

///Response
type PesaLinkResponse struct {
	TransactionId  string `json:"transactionId"`
	Status         string `json:"status"`
	Description    string `json:"description"`
	ResponseCode   string `json:"response_code"`
	ResponseMsg    string `json:"response_msg"`
	ResponseStatus string `json:"response_status"`
}

////PesaLink ends here

type SendMoneyResponse struct {
	TransactionID  string `json:"transactionId"`
	Status         string `json:"status"`
	ResponseStatus string `json:"response_status"`
	ResponseMsg    string `json:"response_msg"`
	ResponseCode   string `json:"response_code"`
}
