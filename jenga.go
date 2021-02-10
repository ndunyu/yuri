package yuri

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type Jenga struct {
	Live           bool
	Username       string
	Password       string
	MerchantCode   string
	ApiKey         string
	PrivateKeyPath string
}

func NewJenga(Username, Password, ApiKey, MerchantCode, PrivateKeyPath string, live bool) Jenga {
	return Jenga{
		Live:         live,
		MerchantCode: MerchantCode,

		Password:       Password,
		Username:       Username,
		ApiKey:         ApiKey,
		PrivateKeyPath: PrivateKeyPath,
	}

}

///https://sandbox.jengahq.io/account-test/v2/accounts/balances/countryCode/accountId

func (J *Jenga) GetEazzyPayMerchants(page, per_page string) (map[string]interface{}, error) {
	queryParameters := make(map[string]string)
	queryParameters["page"] = page
	queryParameters["per_page"] = per_page
	merchants := make(map[string]interface{})
	err := J.getAndProcessJengaRequest(J.getJengaMerchantsUrl(), "", &merchants, queryParameters, nil)

	return merchants, err

}
func (J *Jenga) GetAccountBalance(countryCode, accountId string) (map[string]interface{}, error) {
	merchants := make(map[string]interface{})
	sigString := joinStrings(countryCode, accountId)

	err := J.getAndProcessJengaRequest(J.getAccountBalanceUrl(countryCode, accountId), sigString, &merchants, nil, nil)
	return merchants, err
}

func (J *Jenga) BankToMobileMoneyTransfer(request BankToMobileMoneyRequest) (*SendMoneyResponse, error) {
	var sendMoneyResponse SendMoneyResponse
	request.Destination.Type = "mobile"
	request.Transfer.Type = "MobileWallet"
	var sigString string
	if request.Destination.WalletName == Equitel {
		sigString = joinStrings(request.Source.AccountNumber, request.Transfer.Amount, request.Transfer.CurrencyCode, request.Transfer.Reference)

	} else {
		sigString = joinStrings(request.Transfer.Amount, request.Transfer.CurrencyCode, request.Transfer.Reference, request.Source.AccountNumber)

	}

	err := J.sendAndProcessJengaRequest(J.getBankToMobileWalletUrl(), sigString, request, &sendMoneyResponse, nil)
	return &sendMoneyResponse, err

}
func (J *Jenga) PesaLinkMoneyTransfer(request PesaLinkRequest) (*PesaLinkResponse, error) {
	var pesaLinkResponse PesaLinkResponse
	///var m map[string]interface{}
	request.Destination.Type = "bank"
	request.Transfer.Type = "PesaLink"
	var sigString string

	sigString = joinStrings( request.Transfer.Amount,request.Transfer.CurrencyCode, request.Transfer.Reference, request.Destination.Name, request.Source.AccountNumber)
     log.Println(sigString)
	log.Println("Item are pelob here")
	err := J.sendAndProcessJengaRequest(J.getPesaLinkToBankUrl(), sigString, request, &pesaLinkResponse, nil)
	return &pesaLinkResponse, err
}
func (J *Jenga) EquityToEquityMoneyTransfer(request PesaLinkRequest) (*PesaLinkResponse, error) {
	request.Destination.Type = "bank"
	var pesaLinkResponse PesaLinkResponse
	request.Transfer.Type="InternalFundsTransfer"
	sigString := joinStrings( request.Source.AccountNumber,request.Transfer.Amount, request.Transfer.CurrencyCode, request.Transfer.Reference)
	err := J.sendAndProcessJengaRequest(J.getEquityToEquityUrl(), sigString, request, &pesaLinkResponse, nil)
	return &pesaLinkResponse, err

}


func (J *Jenga) PurchaseAirtime(airtimeRequest AirtimeRequest) (*AirtimeResponse, error) {

	var airTimeResponse AirtimeResponse
	sigString := joinStrings(J.MerchantCode, airtimeRequest.Airtime.Telco, airtimeRequest.Airtime.Amount, airtimeRequest.Airtime.Reference) /// J.MerchantCode + airtimeRequest.Airtime.Telco + airtimeRequest.Airtime.Amount + airtimeRequest.Airtime.Reference

	err := J.sendAndProcessJengaRequest(J.getAirTimeUrl(), sigString, airtimeRequest, &airTimeResponse, nil)
	return &airTimeResponse, err

}

//will verify users
//National iD number
func (J *Jenga) VerifyUserKyc(identityRequestBody IdentityRequestBody) (*IdentityResponseBody, error) {

	var identityResponseBody IdentityResponseBody
	sigString := joinStrings(J.MerchantCode, identityRequestBody.Identity.DocumentNumber, identityRequestBody.Identity.CountryCode)
	err := J.sendAndProcessJengaRequest(J.getKycUrl(), sigString, identityRequestBody, &identityResponseBody, nil)
	return &identityResponseBody, err
}

func joinStrings(items ...string) string {
	var joinedString string
	for _, item := range items {
		joinedString = joinedString + item
	}
	return joinedString
}
func (J *Jenga) GetAccessToken() (*JengaAccessToken, error) {
	//"https://uat.jengahq.io/identity/v2/token"
	data := url.Values{}
	data.Set("username", J.Username)
	data.Set("password", J.Password)
	requestBody := strings.NewReader(data.Encode())
	req, err := http.NewRequest(http.MethodPost, J.getAccessTokenUrl(), requestBody)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+J.ApiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		///Timeout: 15 * time.Second
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}

	var token JengaAccessToken
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {

		return nil, errors.New("error converting from json")
	}
	return &token, nil

}

func (J *Jenga) getAndProcessJengaRequest(url, sigString string, response interface{}, queryParameters, extraHeader map[string]string) error {
	if reflect.ValueOf(response).Kind() != reflect.Ptr {
		log.Println("not a pointer")

		return errors.New("response should be a pointer")

	}
	token, err := J.GetAccessToken()
	if err != nil {

		return err
	}


	headers := make(map[string]string)
	if !IsEmpty(sigString) {
		s := strings.TrimSpace(sigString)
		signature, err := SignSha256DataWithPrivateKey(s, J.PrivateKeyPath)
		if err != nil {

			return err
		}

		headers["signature"] = signature

	}

	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token.AccessToken
	for k, v := range extraHeader {
		headers[k] = v
	}
	resp, err := getRequest(url, headers, queryParameters)
	if err != nil {

		return err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		log.Println(resp.StatusCode)
		b, _ := ioutil.ReadAll(resp.Body)

		return &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}

	///var dt map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {

		return errors.New("error converting from json")
	}

	return nil
}

//make sure response is a pointer
func (J *Jenga) sendAndProcessJengaRequest(url, sigString string, data interface{}, response interface{}, extraHeader map[string]string) error {

	if reflect.ValueOf(response).Kind() != reflect.Ptr {
		log.Println("not a pointer")

		return errors.New("response should be a pointer")

	}
	token, err := J.GetAccessToken()
	if err != nil {
		return err
	}
	s := strings.TrimSpace(sigString)
	signature, err := SignSha256DataWithPrivateKey(s, J.PrivateKeyPath)
	if err != nil {

		return err
	}
	headers := make(map[string]string)

	headers["signature"] = signature
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token.AccessToken
	for k, v := range extraHeader {
		headers[k] = v
	}

	resp, err := postRequest(url, data, headers)
	if err != nil {

		return err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {

		b, _ := ioutil.ReadAll(resp.Body)

		return &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}


	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {

		return errors.New("error converting from json")
	}

	return nil
}
