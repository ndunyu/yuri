package yuri

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Mpesa struct {
	Live           bool
	ConsumerKey    string
	ConsumerSecret string
}

func NewMpesa(ConsumerKey, ConsumerSecret string, live bool) Mpesa {
	return Mpesa{
		Live:           live,
		ConsumerKey:    ConsumerKey,
		ConsumerSecret: ConsumerSecret,
	}

}

func (m *Mpesa) SetMode(mode bool) {
	m.Live = mode

}

//GetAccessToken will get the token to be used to query data
func (m *Mpesa) GetAccessToken() (*AccessTokenResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.getAccessTokenUrl(), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(m.ConsumerKey, m.ConsumerSecret)

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}
	var token AccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {

		return nil, errors.New("error converting from json")
	}
	return &token, nil
}

//B2C Sends Money from a business to the Customer
func (m *Mpesa) B2CRequest(b2c B2CRequestBody) (*MpesaResult, error) {
	token, err := m.GetAccessToken()
	if err != nil {

		return nil, err
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token.AccessToken
	//url:="https://sandbox.safaricom.co.ke/mpesa/b2c/v1/paymentrequest"
	url := m.getB2CUrl()

	resp, err := postRequest(url, b2c, headers)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}
	var response MpesaResult

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {

		PrintStruct(err)
		return nil, errors.New("error converting from json")
	}
	return &response, nil
}

//B2C Sends Money from a business to the Customer
func (m *Mpesa) B2BRequest(b2b B2BRequestBody) (*MpesaResult, error) {
	token, err := m.GetAccessToken()
	if err != nil {

		return nil, err
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token.AccessToken
	//url:="https://sandbox.safaricom.co.ke/mpesa/b2c/v1/paymentrequest"
	url := m.getB2BUrl()

	resp, err := postRequest(url, b2b, headers)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}
	var response MpesaResult

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {

		PrintStruct(err)
		return nil, errors.New("error converting from json")
	}
	return &response, nil

}



//B2C Sends Money from a business to the Customer
func (m *Mpesa) C2BRequest(b2b B2BRequestBody) (*MpesaResult, error) {


	return nil,nil

}



func (m *Mpesa)AccountBalanceRequest (balance AccountBalanceRequestBody)(*MpesaResult, error) {
	token, err := m.GetAccessToken()
	if err != nil {

		return nil, err
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token.AccessToken
	//url:="https://sandbox.safaricom.co.ke/mpesa/b2c/v1/paymentrequest"
	url := m.getBalanceUrl()

	resp, err := postRequest(url, balance, headers)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}
	var response MpesaResult

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {

		PrintStruct(err)
		return nil, errors.New("error converting from json")
	}


	return &response, nil
}


func getRequest(url string, headers map[string]string) (*http.Response, error) {
	///requestBody, err := json.Marshal(data)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 15 * time.Second}
	return client.Do(req)

}

func postRequest(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	///requestBody, err := json.Marshal(data)

	b, err := json.Marshal(data)
	if err != nil {

		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{Timeout: 15 * time.Second}
	return client.Do(req)

}

func (m *Mpesa) GetSecurityCredential(initiatorPassword string) (string, error) {
	var fileName string
	if m.Live {
		fileName = "https://developer.safaricom.co.ke/sites/default/files/cert/cert_prod/cert.cer"
	} else {
		fileName = "https://developer.safaricom.co.ke/sites/default/files/cert/cert_sandbox/cert.cer"
		log.Println("herre")

	}
	resp, err := getRequest(fileName, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var key []byte
	key, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(key)
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	rng := rand.Reader

	encrypted, err := rsa.EncryptPKCS1v15(rng, rsaPublicKey, []byte(initiatorPassword))
	if err != nil {
		return "", err
	}

	enc := base64.StdEncoding.EncodeToString(encrypted)
	println(enc)
	return enc, nil

}

