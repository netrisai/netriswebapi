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

package port

import (
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type PortClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *PortClient {
	return &PortClient{c}
}

func parse(APIResult *http.APIResponse) ([]*Port, error) {
	var items []*Port
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parseAPIVnet} %s", err)
	}
	return items, nil
}

func (c *PortClient) Get() ([]*Port, error) {
	address := c.client.URL.String() + v2address.Ports
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}