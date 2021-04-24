package yuri

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const LiveUrl = "https://api.africastalking.com/version1/"
const SandboxUrl = "https://api.sandbox.africastalking.com/version1/"
const SMSUrl = "messaging"
const AirtimeUrl = "airtime/send"


type AfricaTalking struct {
	Live     bool
	ApiKey   string
	UserName string
	From     string
}

func NewAfricaTalking(apiKey, userName, from string) AfricaTalking {
	return AfricaTalking{
		Live:     true,
		ApiKey:   apiKey,
		UserName: userName,
		From:     from,
	}

}

func (a *AfricaTalking) SetMode(mode bool) {

	a.Live = mode

}

func (a *AfricaTalking) SendSms(to, message string) (*AfricaTalkingResponse, error) {

	body := url.Values{}
	body.Set("username", a.UserName)
	body.Set("to", to)
	body.Set("message", message)
	body.Set("from", a.From)
	resp, err := a.postRequest(a.getSmsUrl(), body, nil)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	var response AfricaTalkingResponse



	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {


		return nil, errors.New("error converting from json")
	}
	return &response, nil
}
func (a *AfricaTalking) SendAirtime(airtimeRequest []AfricaTalkingAirtimeRequest) (*AfricaTalkingMessageResponse, error) {


	africaTalkingMaps := []map[string]string{}
	for i, _ := range airtimeRequest {
	   dataMap:=airtimeRequest[i].ToAfricaTalkingString()
	   africaTalkingMaps=append(africaTalkingMaps,dataMap)

	}
	data,err:=ToJson(africaTalkingMaps)
	if err!=nil {
		return nil, err
	}
	africaTalkingString:=string(data)



	body := url.Values{}
	body.Set("username", a.UserName)
	body.Set("recipients", africaTalkingString)
	resp, err := a.postRequest(a.getAirtimeUrl(), body, nil)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	var response AfricaTalkingMessageResponse



	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}

	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {

		return nil, errors.New("error converting from json")
	}
	return &response, nil


}





func (a *AfricaTalking) getBaseUrl() string {
	if !a.Live {

		return SandboxUrl
	}

	return LiveUrl
}



func (a *AfricaTalking) getSmsUrl() string {

	return  a.getBaseUrl()+SMSUrl

}
func (a *AfricaTalking) getAirtimeUrl() string {

	return  a.getBaseUrl()+AirtimeUrl

}

func (a *AfricaTalking) postRequest(url string, data url.Values, headers map[string]string) (*http.Response, error) {
	///requestBody, err := json.Marshal(data)

	requestBody := strings.NewReader(data.Encode())

	req, err := http.NewRequest(http.MethodPost, url, requestBody)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("apikey", a.ApiKey)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{
		///Timeout: 15 * time.Second
	}
	return client.Do(req)

}

type AfricaTalkingResponse struct {
	SMSMessageData SMSMessageData `json:"SMSMessageData"`
}

type Recipients struct {
	Number string `json:"number"`

	Status     string `json:"status"`
	Cost       string `json:"cost"`
	MessageId  string `json:"messageId"`
	StatusCode int    `json:"statusCode"`
}

type SMSMessageData struct {
	Message string `json:"Message"`

	Recipients []Recipients `json:"Recipients"`
}

type AfricaTalkingAirtimeRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Currency   string `json:"currency"`
	Amount float64 `json:"amount"`

}
type AfricaTalkingMessageResponse struct {
	ErrorMessage  string      `json:"errorMessage"`
	NumSent       int         `json:"numSent"`
	TotalAmount   string      `json:"totalAmount"`
	TotalDiscount string      `json:"totalDiscount"`
	Responses     []Responses `json:"responses"`
}
type Responses struct {
	PhoneNumber  string `json:"phoneNumber"`
	ErrorMessage string `json:"errorMessage"`
	Amount       string `json:"amount"`
	Status       string `json:"status"`
	RequestID    string `json:"requestId"`
	Discount     string `json:"discount"`
}



func (a AfricaTalkingAirtimeRequest) ToAfricaTalkingString() map[string]string {
	data := map[string]string{
		"phoneNumber": a.PhoneNumber,
		"amount":      a.Currency + " " + fmt.Sprintf("%.2f", a.Amount) ,
	}
	return  data


}


type RequestError struct {
	StatusCode int

	Message string
	Url string
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("url is: %s \n status code is: %d \n  and body is : %s",r.Url, r.StatusCode, r.Message)


}
