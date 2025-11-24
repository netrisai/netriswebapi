/*
Copyright 2023. Netris, Inc.

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

package nvlinkledger

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type Client struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *Client {
	return &Client{c}
}

func parse(APIResult *http.APIResponse) ([]*NVLinkledger, error) {
	var items []*NVLinkledger
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

// nolint
func parseSingle(APIResult *http.APIResponse) (*NVLinkledger, error) {
	var items *NVLinkledger
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func (c *Client) Get() ([]*NVLinkledger, error) {
	address := c.client.URL.String() + v2address.NVLinkledger
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetNVLinkledger} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetNVLinkledger} %s", err)
	}
	return items, nil
}

func (c *Client) GetByNmxcID(nmxcID string) ([]*NVLinkledger, error) {
	address := c.client.URL.String() + v2address.NVLinkledger + fmt.Sprintf("/?nmxcID=%s", nmxcID)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetNVLinkledgerByNmxcID} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetNVLinkledgerByNmxcID} %s", err)
	}
	return items, nil
}

func (c *Client) GetByNmxcIDAndPartitionID(nmxcID string, partitionID int) ([]*NVLinkledger, error) {
	address := c.client.URL.String() + v2address.NVLinkledger + fmt.Sprintf("?nmxcID=%s&partitionID=%d", nmxcID, partitionID)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetNVLinkledgerByNmxcIDAndPartitionID} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetNVLinkledgerByNmxcIDAndPartitionID} %s", err)
	}
	return items, nil
}

func (c *Client) Add(nvlinkledger *NVLinkledgerW) (reply http.HTTPReply, err error) {
	if nvlinkledger != nil && nvlinkledger.Custom == nil {
		nvlinkledger.Custom = map[string]interface{}{}
	}
	js, err := json.Marshal(nvlinkledger)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.NVLinkledger
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) Update(nmxcID string, partitionID int, nvlinkledger *NVLinkledgerU) (reply http.HTTPReply, err error) {
	if nvlinkledger != nil && nvlinkledger.Custom == nil {
		nvlinkledger.Custom = map[string]interface{}{}
	}
	js, err := json.Marshal(nvlinkledger)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdateNVLinkledger} %s", err)
	}
	address := c.client.URL.String() + v2address.NVLinkledger + fmt.Sprintf("/%s/%d", nmxcID, partitionID)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdateNVLinkledger} %s", err)
	}

	return reply, nil
}

func (c *Client) Delete(nmxcID string, partitionID int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.NVLinkledger + fmt.Sprintf("/%s/%d", nmxcID, partitionID)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
