package yuri

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Jenga struct {
	Live         bool
	Username     string
	Password     string
	MerchantCode string
	ApiKey       string
	PrivateKeyPath string
}

func NewJenga(Username, Password, ApiKey,MerchantCode ,PrivateKeyPath string, live bool) Jenga {
	return Jenga{
		Live: live,
		MerchantCode: MerchantCode,

		Password: Password,
		Username: Username,
		ApiKey:   ApiKey,
		PrivateKeyPath: PrivateKeyPath,

	}

}

//will verify users
//National iD number
func (J *Jenga) VerifyUserKyc(identityRequestBody IdentityRequestBody) (*IdentityResponseBody, error) {
	token, err := J.GetAccessToken()
	if err != nil {
		return nil, err
	}
	sigString := J.MerchantCode + identityRequestBody.Identity.DocumentNumber + identityRequestBody.Identity.CountryCode
	signature, err := SignSha256DataWithPrivateKey(sigString, J.PrivateKeyPath)
	if err != nil {

		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token.AccessToken
	headers["signature"] = signature

	resp, err := postRequest(J.getKycUrl(), identityRequestBody, headers)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}
	var identityResponseBody IdentityResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&identityResponseBody); err != nil {

		return nil, errors.New("error converting from json")
	}
	return &identityResponseBody, nil

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
