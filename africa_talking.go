package yuri

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const LiveUrl = "https://api.africastalking.com/version1/messaging"
const SandboxUrl = "https://api.sandbox.africastalking.com/version1/messaging"

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
	resp, err := a.postRequest(a.getUrl(), body, nil)
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

func (a *AfricaTalking) getUrl() string {
	if !a.Live {

		return SandboxUrl
	}

	return LiveUrl
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

type RequestError struct {
	StatusCode int

	Message string
}

func (r *RequestError) Error() string {
	return r.Message
}
