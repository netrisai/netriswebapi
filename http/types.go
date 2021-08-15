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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

// HTTPCred stores the credentials for connect to http server. User, Password, HTTP Clien e.t.c
type HTTPCred struct {
	sync.Mutex
	URL                url.URL
	LoginData          loginData
	Cookies            []http.Cookie
	ConnectSID         string
	Timeout            int
	InsecureSkipVerify bool
	RedirectCount      int
}

// HTTPReply represents the structure of the response which has come from Netris API
type HTTPReply struct {
	Data       []byte
	StatusCode int
	Status     string
	Error      error
}

type loginData struct {
	Login      string `json:"user"`
	Password   string `json:"password"`
	AuthScheme int    `json:"auth_scheme_id"`
}

func (r *HTTPReply) Parse() (*APIResponse, error) {
	var result *APIResponse
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		return nil, fmt.Errorf("{parseAPIResponse} %s", err)
	}
	return result, nil
}

func (r *HTTPReply) String() (string, error) {
	js, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(js), err
}

// APIResponse .
type APIResponse struct {
	Data      map[string]interface{} `json:"data"`
	Errors    APIResponseErrors      `json:"errors"`
	Message   string                 `json:"message"`
	IsSuccess bool                   `json:"isSuccess"`
	Meta      APIResponseMeta        `json:"meta"`
}

// APIResponseErrors .
type APIResponseErrors struct {
	Name  string `json:"name"`
	Error string `json:"error"`
}

// APIResponseMessage .
type APIResponseMessage struct {
	Message string `json:"message"`
}

// APIResponseMeta .
type APIResponseMeta struct {
	APIVersion string `json:"apiVersion"`
	StatusCode int    `json:"statusCode"`
}
