package airtime

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AndroidStudyOpenSource/africastalking-go/util"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Response is the reponse from the api
type Response struct {
	NumSent       int
	TotalAmount   string
	TotalDiscount string
	ErrorMessage  string
	Responses     []Entry
}

// Entry is the entry for each airtime response
type Entry struct {
	ErrorMessage string
	PhoneNumber  string
	Amount       string
	Discount     string
	Status       string
	RequestID    string
}

// Service is the airtime service
type Service struct {
	Username string
	APIKey   string
	Env      string
}

// NewService returns a new service
func NewService(username, apiKey, env string) Service {
	return Service{username, apiKey, env}
}

// Send sends a new airtime request
func (service Service) Send() (*Response, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/version1/airtime"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request %v", err)
	}

	values := request.URL.Query()
	values.Add("username", service.Username)
	request.URL.RawQuery = values.Encode()

	request.Header.Set("apikey", service.APIKey)
	request.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("could not get rsponse %v", err)
	}
	defer response.Body.Close()

	var airtimeResponse Response
	json.NewDecoder(response.Body).Decode(&airtimeResponse)
	return &airtimeResponse, nil
}

func (service Service) SendAirtime(phoneNumber, amount string) (*Response, error) {
	values := url.Values{}
	values.Set("username", service.Username)
	values.Set("phoneNumber", phoneNumber)
	values.Set("amount", amount)

	smsURL := util.GetSmsURL(service.Env)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	res, err := service.newPostRequest(smsURL, values, headers)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, errors.New("unable to parse sms response")
	}
	return &response, nil
}

func (service Service) newPostRequest(url string, values url.Values, headers map[string]string) (*http.Response, error) {
	reader := strings.NewReader(values.Encode())

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Length", strconv.Itoa(reader.Len()))
	req.Header.Set("apikey", service.APIKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}
