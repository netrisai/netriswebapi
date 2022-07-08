/*
Copyright 2021. Netris, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
}

// NewHTTPCredentials .
func NewHTTPCredentials(address, login, password string, timeout int) (*HTTPCred, error) {
	URL, err := url.Parse(strings.TrimSuffix(address, "/"))
	if err != nil {
		return nil, fmt.Errorf("{newHTTPCredentials} %s", err)
	}
	if timeout == 0 {
		timeout = 5
	}
	return &HTTPCred{
		URL: *URL,
		LoginData: loginData{
			Login:    login,
			Password: password,
		},
		Timeout:       timeout,
		RedirectCount: 10,
	}, nil
}

// NewHTTPCredentialsWithCooke .
func NewHTTPCredentialsWithCooke(address, sessionID string, timeout int) (*HTTPCred, error) {
	URL, err := url.Parse(strings.TrimSuffix(address, "/"))
	if err != nil {
		return nil, fmt.Errorf("{NewHTTPCredentialsWithCooke} %s", err)
	}
	if timeout == 0 {
		timeout = 5
	}
	return &HTTPCred{
		URL:           *URL,
		LoginData:     loginData{},
		Cookies:       []http.Cookie{{Name: "connect.sid", Value: sessionID}},
		Timeout:       timeout,
		RedirectCount: 10,
	}, nil
}

// InsecureVerify .
func (cred *HTTPCred) InsecureVerify(enable bool) {
	cred.Lock()
	defer cred.Unlock()
	cred.InsecureSkipVerify = enable
}

// CheckAuth checks the user authorized or not
func (cred *HTTPCred) CheckAuth() error {
	reply, err := cred.GetRequest(cred.URL.String() + v2address.Auth)
	if err != nil {
		return fmt.Errorf("{CheckAuth} %s", err)
	}
	if reply.StatusCode == http.StatusOK {
		return nil
	} else if len(reply.Data) > 0 {
		return fmt.Errorf("{CheckAuth} %s", reply.Data)
	}
	return fmt.Errorf("{CheckAuth} not authorized")
}

func (cred *HTTPCred) SetCookie(sessionID string) {
	cred.setCookie(sessionID)
}

func (cred *HTTPCred) setCookie(sessionID string) {
	cred.Lock()
	defer cred.Unlock()
	cred.Cookies = []http.Cookie{{Name: "connect.sid", Value: sessionID}}
}

func (cred *HTTPCred) loginUser(URL string, redirectCounter int) error {
	reqData := fmt.Sprintf(`{"user": "%s", "password": "%s", "auth_scheme_id": 1}`, cred.LoginData.Login, cred.LoginData.Password)

	req, err := http.NewRequest("POST", URL, bytes.NewBufferString(reqData))
	if err != nil {
		return fmt.Errorf("{LoginUser} %s", err)
	}

	req.Header.Set("Content-type", "application/json")

	client := cred.createHTTPClient()

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("{LoginUser} %s", err)
	}

	if resp.StatusCode == 301 && redirectCounter > 0 {
		location, _ := resp.Location()
		return cred.loginUser(location.String(), redirectCounter-1)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("{LoginUser} Authentication failed")
	}

	cred.Lock()
	var cookies []http.Cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "connect.sid" {
			cred.ConnectSID = cookie.Value
		}
		cookies = append(cookies, *cookie)
	}
	cred.Cookies = cookies
	cred.Unlock()
	return nil
}

// LoginUser login the user and get the cookies for future use
func (cred *HTTPCred) LoginUser() error {
	URL := cred.URL.String() + v2address.Auth
	return cred.loginUser(URL, cred.RedirectCount)
}

// CustomBodyRequest impelements the POST, PUT, UPDATE requests
func (cred *HTTPCred) CustomBodyRequest(method string, address string, data []byte) (reply HTTPReply, err error) {
	return cred.customBodyRequest(method, address, data, cred.RedirectCount)
}

func (cred *HTTPCred) customBodyRequest(method string, address string, data []byte, redirectCounter int) (reply HTTPReply, err error) {
	requestBody := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, address, requestBody)
	if err != nil {
		return reply, fmt.Errorf("{CustomBodyRequest} [%s] [%s] %s", method, address, err)
	}

	request.Header.Set("Content-type", "application/json")
	cred.Lock()
	for _, cookie := range cred.Cookies {
		cook := cookie
		request.AddCookie(&cook)
	}
	client := cred.createHTTPClient()
	cred.Unlock()

	resp, err := client.Do(request)
	if err != nil {
		return reply, fmt.Errorf("{CustomBodyRequest} [%s] [%s] %s", method, address, err)
	}

	if resp.StatusCode == 301 && redirectCounter > 0 {
		location, _ := resp.Location()
		return cred.customBodyRequest(method, location.String(), data, redirectCounter-1)
	}

	reply.StatusCode = resp.StatusCode
	reply.Status = resp.Status

	defer resp.Body.Close()
	reply.Data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return reply, fmt.Errorf("{CustomBodyRequest} [%s] [%s] %s", method, address, err)
	}

	return reply, nil
}

// Post custom request
func (cred *HTTPCred) Post(address string, data []byte) (reply HTTPReply, err error) {
	return cred.CustomBodyRequest("POST", address, data)
}

// Put custom request
func (cred *HTTPCred) Put(address string, data []byte) (reply HTTPReply, err error) {
	return cred.CustomBodyRequest("PUT", address, data)
}

// Delete custom request
func (cred *HTTPCred) Delete(address string, data []byte) (reply HTTPReply, err error) {
	return cred.CustomBodyRequest("DELETE", address, data)
}

// ParseAPIResponse .
func ParseAPIResponse(data []byte) (*APIResponse, error) {
	var result *APIResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("{parseAPIResponse} %s", err)
	}
	return result, nil
}

// Decode .
func Decode(input interface{}, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return fmt.Errorf("{customDecode} %s", err)
	}

	err = decoder.Decode(input)
	if err != nil {
		return fmt.Errorf("{customDecode} %s", err)
	}
	return err
}

// Get .
func (cred *HTTPCred) Get(address string) (*APIResponse, error) {
	data, err := cred.GetRequest(address)
	if err != nil {
		return nil, fmt.Errorf("{http.Get} %s", err)
	}

	if data.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("{http.Get} %s", data.Data)
	}

	result, err := ParseAPIResponse(data.Data)
	if err != nil {
		return nil, fmt.Errorf("{http.Get} %s", err)
	}
	return result, nil
}

// GetRequest .
func (cred *HTTPCred) GetRequest(address string) (reply HTTPReply, err error) {
	return cred.getRequest(address, cred.RedirectCount)
}

func (cred *HTTPCred) getRequest(address string, redirectCounter int) (reply HTTPReply, err error) {
	request, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return reply, fmt.Errorf("{http.get} [%s] %s", address, err)
	}

	request.Header.Set("Content-type", "application/json")
	cred.Lock()
	for _, cookie := range cred.Cookies {
		cook := cookie
		request.AddCookie(&cook)
	}
	client := cred.createHTTPClient()
	cred.Unlock()

	resp, err := client.Do(request)
	if err != nil {
		return reply, fmt.Errorf("{http.get} [%s] %s", address, err)
	}

	if resp.StatusCode == 301 && redirectCounter > 0 {
		location, _ := resp.Location()
		return cred.getRequest(location.String(), redirectCounter-1)
	}

	reply.StatusCode = resp.StatusCode
	reply.Status = resp.Status

	defer resp.Body.Close()
	reply.Data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return reply, fmt.Errorf("{http.get} [%s] %s", address, err)
	}
	return reply, err
}

func (cred *HTTPCred) createHTTPClient() http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: cred.InsecureSkipVerify},
	}
	return http.Client{
		Timeout:   time.Duration(time.Duration(cred.Timeout) * time.Second),
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

// CheckAuthWithInterval .
func (cred *HTTPCred) CheckAuthWithInterval() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		err := cred.CheckAuth()
		if err != nil {
			logger.Error(err)
			err := cred.LoginUser()
			if err != nil {
				logger.Error(err.Error())
			}
		}
	}
}
