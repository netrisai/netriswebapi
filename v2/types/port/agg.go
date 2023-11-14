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

func (c *PortClient) GetAggregatedPorts() ([]*AggregatedPort, error) {
	address := c.client.URL.String() + v2address.PortAggregated
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetAggregatedPorts} %s", err)
	}

	items, err := parseAggList(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetAggregatedPorts} %s", err)
	}
	return items, nil
}

func (c *PortClient) GetAggregatedPortByID(id int) (*AggregatedPort, error) {
	ports, err := c.GetAggregatedPorts()
	if err != nil {
		return nil, fmt.Errorf("{GetAggregatedPortByID} %s", err)
	}
	for _, p := range ports {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, fmt.Errorf("{GetAggregatedPortByID} Port with id %d not found", id)
}

func parseAggList(APIResult *http.APIResponse) ([]*AggregatedPort, error) {
	var items []*AggregatedPort
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parseAggList} %s", err)
	}
	return items, nil
}
