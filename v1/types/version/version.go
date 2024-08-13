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

package version

import (
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
)

type VersionClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *VersionClient {
	return &VersionClient{c}
}

func parse(APIResult *http.APIResponse) (Version, error) {
	var version Version
	err := http.Decode(APIResult.Data, &version)
	if err != nil {
		return version, fmt.Errorf("{parse} %s", err)
	}
	return version, nil
}

func (c *VersionClient) Get() (Version, error) {
	address := c.client.URL.String() + v1address.Version
	APIResult, err := c.client.Get(address)
	var responce Version
	if err != nil {
		return responce, fmt.Errorf("{GetVersion} %s", err)
	}

	responce, err = parse(APIResult)
	if err != nil {
		return responce, fmt.Errorf("{GetVersion} %s", err)
	}
	return responce, nil
}
