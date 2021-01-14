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
	"time"
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

func (J *Jenga) BankToMobileMoneyTransfer(bankToMobileMoneyRequest BankToMobileMoneyRequest) {

}

func (J *Jenga) PurchaseAirtime(airtimeRequest AirtimeRequest) (*AirtimeResponse, error) {

	var airTimeResponse AirtimeResponse
	sigString := J.MerchantCode + airtimeRequest.Airtime.Telco + airtimeRequest.Airtime.Amount + airtimeRequest.Airtime.Reference
	signature, err := SignSha256DataWithPrivateKey(sigString, J.PrivateKeyPath)
	if err != nil {

		return nil, err
	}
	headers := make(map[string]string)
	headers["signature"] = signature
	err = J.sendAndProcessJengaRequest(J.getAirTimeUrl(), airtimeRequest, &airTimeResponse, headers)
	return &airTimeResponse, err

}

//will verify users
//National iD number
func (J *Jenga) VerifyUserKyc(identityRequestBody IdentityRequestBody) (*IdentityResponseBody, error) {

	var identityResponseBody IdentityResponseBody
	sigString := J.MerchantCode + identityRequestBody.Identity.DocumentNumber + identityRequestBody.Identity.CountryCode
	signature, err := SignSha256DataWithPrivateKey(sigString, J.PrivateKeyPath)
	if err != nil {

		return nil, err
	}

	headers := make(map[string]string)
	headers["signature"] = signature
	err = J.sendAndProcessJengaRequest(J.getKycUrl(), identityRequestBody, &identityResponseBody, headers)
	return &identityResponseBody, err
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
	client := &http.Client{Timeout: 15 * time.Second}
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

//make sure response is a pointer
func (J *Jenga) sendAndProcessJengaRequest(url string, data interface{}, response interface{}, extraHeader map[string]string) error {

	if reflect.ValueOf(response).Kind() != reflect.Ptr {
		log.Println("not a pointer")

		return errors.New("response should be a pointer")

	}
	token, err := J.GetAccessToken()
	if err != nil {
		return err
	}
	headers := make(map[string]string)
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
