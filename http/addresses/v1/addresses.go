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

package v1

const (

	Sites     = "/api/sites"
	GSettings = "/api/general"
	L4LB      = "/api/l4lb"
	Subnets   = "/api/subnets"

	ACL = "/api/acl"

	ACL2                   = "/api/acltwozero"
	ACL2Status             = "/api/acltwozero/changestatus"
	ACL2Publishers         = "/api/acltwozero/publishers"
	ACL2PublishersApprove  = "/api/acltwozero/publishers/approve"
	ACL2PublishersReject   = "/api/acltwozero/publishers/reject"
	ACL2Subscribers        = "/api/acltwozero/subscribers"
	ACL2SubscribersApprove = "/api/acltwozero/subscribers/approve"
	ACL2SubscribersReject  = "/api/acltwozero/subscribers/reject"

	PortGroup = "/api/aclportgroups"

	GraphBoards = "/api/graphboards"

	Inventory        = "/api/inventory"
	InventoryUsedIPs = "/api/inventory/usedips"

	InventoryProfiles = "/api/inventoryProfiles"

	VNet         = "/api/v-net"
	VNetInfo     = "/api/v-net/info"
	VNetValidate = "/api/v-net/validate"

	Ports   = "/api/switchports"
	Tenants = "/api/tenants"

	Users = "/api/users"

	PermissionGroups = "/api/permissiongroups"

	UserRoles = "/api/userroles"

	BGP               = "/api/ebgp"
	BGPSites          = "/api/ebgp/sites"
	BGPVNets          = "/api/ebgp/v-nets"
	BGPRouteMaps      = "/api/ebgp/routemaps"
	BGPOffloaders     = "/api/ebgp/offloaders"
	BGPPorts          = "/api/ebgp/ports"
	BGPSwitches       = "/api/ebgp/switches"
	BGPUpdatedSources = "/api/ebgp/updatesources"

	BGPObjects = "/api/ebgpobjects"

	Routes = "/api/routes"

	RouteMaps = "/api/ebgproutemaps"
)
