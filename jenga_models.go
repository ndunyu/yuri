package yuri

type JengaAccessToken struct {
	TokenType string `json:"token_type"`
	IssuedAt string `json:"issued_at"`
	ExpiresIn string `json:"expires_in"`
	AccessToken string `json:"access_token"`

}

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