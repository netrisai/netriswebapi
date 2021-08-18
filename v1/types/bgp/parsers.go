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

package bgp

import (
	"fmt"

	"github.com/netrisai/netriswebapi/http"
)

func parse(APIResult *http.APIResponse) ([]*EBGP, error) {
	var items []*EBGP
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parseSites(APIResult *http.APIResponse) ([]*EBGPSite, error) {
	var items []*EBGPSite
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parseVNets(APIResult *http.APIResponse) ([]*EBGPVNet, error) {
	var items []*EBGPVNet
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parseRouteMaps(APIResult *http.APIResponse) ([]*EBGPRouteMap, error) {
	var items []*EBGPRouteMap
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parseOffloaders(APIResult *http.APIResponse) ([]*EBGPOffloader, error) {
	var items []*EBGPOffloader
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parsePorts(APIResult *http.APIResponse) ([]*EBGPPort, error) {
	var items []*EBGPPort
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parseSwitches(APIResult *http.APIResponse) ([]*EBGPSwitch, error) {
	var items []*EBGPSwitch
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parseUpdateSources(APIResult *http.APIResponse) ([]*EBGPUpdatedSource, error) {
	var items []*EBGPUpdatedSource
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}
