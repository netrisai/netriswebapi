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

package v2

const (
	Auth = "/api/auth"

	VNetBase = "/api/v2/vnet"

	InventoryBase       = "/api/v2/hw"
	InventorySwitch     = "/api/v2/hw/switch"
	InventoryController = "/api/v2/hw/controller"
	InventorySoftgate   = "/api/v2/hw/softgate"
	InventoryUsedIPs    = "/api/v2/hw/usedips"
	InventoryProfiles   = "/api/v2/hw/profiles"
	InventorySubnets    = "/api/v2/hw/subnets"
	InventoryNOS        = "/api/v2/hw/nos"

	IPAMBase       = "/api/v2/ipam"
	IPAMSubnets    = "/api/v2/ipam/subnets"
	IPAMHosts      = "/api/v2/ipam/hosts"
	IPAMAllocation = "/api/v2/ipam/allocation"
	IPAMSubnet     = "/api/v2/ipam/subnet"

	Ports = "/api/v2/ports"
)
