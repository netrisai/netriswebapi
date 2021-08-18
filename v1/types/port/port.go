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
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
	"github.com/netrisai/netriswebapi/v1/types/site"
)

type PortClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *PortClient {
	return &PortClient{c}
}

func parse(APIResult *http.APIResponse) ([]*Port, error) {
	var sites []*Port
	err := http.Decode(APIResult.Data, &sites)
	if err != nil {
		return sites, fmt.Errorf("{parse} %s", err)
	}
	return sites, nil
}

func (c *PortClient) Get() ([]*Port, error) {
	sites, err := site.New(c.client).Get()
	if err != nil {
		return nil, err
	}
	siteList := ""
	for _, s := range sites {
		siteList += fmt.Sprintf("selectedSites[]=%d&", s.ID)
	}
	address := c.client.URL.String() + v1address.Ports + "?" + siteList
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
