package yuri

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	errors "errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
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

func (m *Mpesa)StkPushRequest(body StKPushRequestBody, passKey string) (*StkPushResult, error) {
	if body.Timestamp=="" {
		t := time.Now()
		fTime := t.Format("20060102150405")
		body.Timestamp= fTime
		body.Password=GeneratePassword(body.BusinessShortCode,passKey, fTime)
	}
	body.TransactionType=CustomerPayBillOnline
	var stkPushResult StkPushResult
	err:=m.sendAndProcessStkPushRequest(m.getStkPush(),body,&stkPushResult,nil)
	return  &stkPushResult,err
}

func (m *Mpesa)StkPushQuery(body StkPushQueryRequestBody,passKey string)(*StkPushQueryResponseBody, error){

	if body.Timestamp=="" {
		t := time.Now()
		fTime := t.Format("20060102150405")
		body.Timestamp= fTime
		body.Password=GeneratePassword(body.BusinessShortCode,passKey, fTime)
	}

	var stkPushResult StkPushQueryResponseBody
	err:=m.sendAndProcessStkPushRequest(m.getStkPush(),body,&stkPushResult,nil)
	return  &stkPushResult,err




}





// B2CRequest Sends Money from a business to the Customer
func (m *Mpesa) B2CRequest(b2c B2CRequestBody) (*MpesaResult, error) {
	if b2c.CommandID == "" {

		b2c.CommandID = BusinessPayment
	}

	return m.sendAndProcessMpesaRequest(m.getB2CUrl(), b2c, nil)

}

// B2BRequest Sends Money from a business to the business
// I.e from onePayBill to another
func (m *Mpesa) B2BRequest(b2b B2BRequestBody) (*MpesaResult, error) {

	return m.sendAndProcessMpesaRequest(m.getB2BUrl(), b2b, nil)

}

//C2BRequest Sends Money from a business to the Customer
///func (m *Mpesa) C2BRequest(c2 B2BRequestBody) (*MpesaResult, error) {

///return nil, nil

///}

// use AccountBalance as the CommandID
// use AccountBalanceRequestBody.indentifier as 4 if its shortcode
//
func (m *Mpesa) AccountBalanceRequest(balance AccountBalanceRequestBody) (*MpesaResult, error) {
	balance.CommandID = AccountBalance
	//
	return m.sendAndProcessMpesaRequest(m.getBalanceUrl(), balance, nil)

}

func (m *Mpesa) TransactionStatusRequest(transactionStatusRequestBody TransactionStatusRequestBody) (*MpesaResult, error) {
	////1 for user
	///4 for organization (indentifiertype)
	transactionStatusRequestBody.CommandID=TransactionStatusQuery
	return m.sendAndProcessMpesaRequest(m.getTransactionStatusUrl(), transactionStatusRequestBody, nil)

}

func (m *Mpesa) RegisterC2BUrls(C2BURLRequestBody RegisterC2BURLRequestBody) (*MpesaResult, error) {
	return m.sendAndProcessMpesaRequest(m.getC2BRegisterUrl(), C2BURLRequestBody, nil)
}

func (m *Mpesa) SimulateC2BTransaction(c2BSimulationRequestBody C2BSimulationRequestBody) (*MpesaResult, error) {

	return m.sendAndProcessMpesaRequest(m.getC2BSimulationUrl(), c2BSimulationRequestBody, nil)

}

func (m *Mpesa) sendAndProcessMpesaRequest(url string, data interface{}, extraHeader map[string]string) (*MpesaResult, error) {
	token, err := m.GetAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token.AccessToken
	///add the extra headers
	//Get union of the headers
	for k, v := range extraHeader {
		headers[k] = v
	}
	resp, err := postRequest(url, data, headers)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}
	var response MpesaResult
	///var respe map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {

		PrintStruct(err)
		return nil, errors.New("error converting from json")
	}
	///PrintStruct(respe)
	return &response, nil

}

func (m *Mpesa) sendAndProcessStkPushRequest(url string, data interface{},respItem interface{}, extraHeader map[string]string) (error) {
	if reflect.ValueOf(respItem).Kind() != reflect.Ptr {
		log.Println("not a pointer")

		return errors.New("response should be a pointer")

	}

	token, err := m.GetAccessToken()
	if err != nil {
		log.Println("ahhaa")
		return  err
	}
	log.Println(token.AccessToken)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token.AccessToken
	///add the extra headers
	//Get union of the headers
	for k, v := range extraHeader {
		headers[k] = v
	}
	resp, err := postRequest(url, data, headers)
	if err != nil {

		return  err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return  &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}

	///var respe map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(respItem); err != nil {

		PrintStruct(err)
		return  errors.New("error converting from json")
	}
	///PrintStruct(respe)
	return  nil

}


func getRequest(url string, headers map[string]string, queryParameters map[string]string) (*http.Response, error) {
	///requestBody, err := json.Marshal(data)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.URL.String())
	for key, value := range queryParameters {
		q := req.URL.Query()
		q.Add(key, value)
		req.URL.RawQuery = q.Encode()

	}
	fmt.Println(req.URL.String())

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 20 * time.Second}
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



//GetAccessToken will get the token to be used to query data
func (m *Mpesa) GetAccessToken() (*AccessTokenResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.getAccessTokenUrl(), nil)
	if err != nil {
		return nil, err
	}
	log.Println(req.URL.String())

	req.SetBasicAuth(m.ConsumerKey, m.ConsumerSecret)

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		log.Println(resp.StatusCode)

		b, _:= ioutil.ReadAll(resp.Body)
		if string(b)=="" {
			return   nil, &RequestError{Message: "Error getting token", StatusCode: resp.StatusCode}

		}
		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}
	log.Println("passed here")
	var token AccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {

		return nil, errors.New("error converting from json")
	}
	return &token, nil
}
func (m *Mpesa) GetSecurityCredential(initiatorPassword string) (string, error) {
	var fileName string
	if m.Live {
		fileName = "https://developer.safaricom.co.ke/sites/default/files/cert/cert_prod/cert.cer"
	} else {
		fileName = "https://developer.safaricom.co.ke/sites/default/files/cert/cert_sandbox/cert.cer"
		log.Println("herre")

	}
	resp, err := getRequest(fileName, nil, nil)
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
func GeneratePassword(shortCode, passkey, time string) string {
	password := fmt.Sprintf("%s%s%s", shortCode,passkey, time)
	return base64.StdEncoding.EncodeToString([]byte(password))

}