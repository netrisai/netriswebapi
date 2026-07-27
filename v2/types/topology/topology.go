/*
Copyright 2026. Netris, Inc.

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

package topology

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	netrishttp "github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

var ufmIDRE = regexp.MustCompile(`^[A-Za-z0-9_.-]+$`)

type Client struct {
	client *netrishttp.HTTPCred
}

func New(c *netrishttp.HTTPCred) *Client {
	return &Client{c}
}

func validateUFMID(ufmID string) error {
	if !ufmIDRE.MatchString(ufmID) {
		return fmt.Errorf("invalid ufm-id %q: use only letters, digits, %q, %q and %q", ufmID, ".", "-", "_")
	}
	return nil
}

func parseImportResult(APIResult *netrishttp.APIResponse) (*UFMImportResult, error) {
	var item UFMImportResult
	err := netrishttp.Decode(APIResult.Data, &item)
	if err != nil {
		return &item, fmt.Errorf("{parse UFMImportResult} %s", err)
	}
	return &item, nil
}

func parsePathResponse(APIResult *netrishttp.APIResponse) (*UFMPathResponse, error) {
	var item UFMPathResponse
	err := netrishttp.Decode(APIResult.Data, &item)
	if err != nil {
		return &item, fmt.Errorf("{parse UFMPathResponse} %s", err)
	}
	return &item, nil
}

func (c *Client) importUFMTopology(ufmID string, links interface{}) (*UFMImportResult, error) {
	if err := validateUFMID(ufmID); err != nil {
		return nil, fmt.Errorf("{ImportUFMTopology} %s", err)
	}

	js, err := json.Marshal(links)
	if err != nil {
		return nil, fmt.Errorf("{ImportUFMTopology} %s", err)
	}

	query := url.Values{}
	query.Set("ufm-id", ufmID)
	address := c.client.URL.String() + v2address.TopologyImport + "?" + query.Encode()

	reply, err := c.client.Post(address, js)
	if err != nil {
		return nil, fmt.Errorf("{ImportUFMTopology} %s", err)
	}
	if reply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("{ImportUFMTopology} %s", reply.Data)
	}

	APIResult, err := netrishttp.ParseAPIResponse(reply.Data)
	if err != nil {
		return nil, fmt.Errorf("{ImportUFMTopology} %s", err)
	}

	item, err := parseImportResult(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{ImportUFMTopology} %s", err)
	}
	return item, nil
}

func (c *Client) ImportUFMTopology(ufmID string, links []*UFMImportLink) (*UFMImportResult, error) {
	return c.importUFMTopology(ufmID, links)
}

func (c *Client) ImportUFMTopologyRaw(ufmID string, links []UFMImportLinkRaw) (*UFMImportResult, error) {
	return c.importUFMTopology(ufmID, links)
}

func (c *Client) GetUFMPath(from string, to ...string) (*UFMPathResponse, error) {
	if len(to) > 1 {
		return nil, fmt.Errorf("{GetUFMPath} expected at most one destination, got %d", len(to))
	}

	query := url.Values{}
	query.Set("from", from)
	if len(to) == 1 {
		query.Set("to", to[0])
	}
	address := c.client.URL.String() + v2address.TopologyPath + "?" + query.Encode()

	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetUFMPath} %s", err)
	}

	item, err := parsePathResponse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetUFMPath} %s", err)
	}
	return item, nil
}
