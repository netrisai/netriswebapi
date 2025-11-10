# Netris Web API Go SDK

Official Go client library for the Netris network automation platform. This SDK provides programmatic access to manage network infrastructure, VPCs, IPAM, Load Balancers, BGP, Server Clusters and more through the Netris API.

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.18-blue)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Authentication](#authentication)
- [Available Resources](#available-resources)
- [Usage Examples](#usage-examples)
- [API Versions](#api-versions)
- [Error Handling](#error-handling)
- [Best Practices](#best-practices)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Comprehensive Resource Management**: Full CRUD operations for network resources
- **Two API Versions**: Support for both V1 and V2 APIs
- **Type-Safe**: Strongly typed Go structs for all resources
- **Session Management**: Automatic cookie-based authentication
- **Flexible Configuration**: Support for custom timeouts and SSL settings
- **Clean API**: Consistent patterns across all resource types

## Installation

```bash
# Install V2 API (recommended)
go get github.com/netrisai/netriswebapi/v2

# Or install V1 API
go get github.com/netrisai/netriswebapi/v1
```

### Requirements

- Go 1.18 or higher
- Network access to Netris controller
- Valid Netris credentials

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    netris "github.com/netrisai/netriswebapi/v2"
)

func main() {
    // Create client
    client, err := netris.Client(
        "https://netris.example.com",
        "username",
        "password",
        30, // timeout in seconds
    )
    if err != nil {
        log.Fatal(err)
    }

    // Authenticate
    if err := client.Client.LoginUser(); err != nil {
        log.Fatal(err)
    }

    // List all sites
    siteClient := client.Site()
    sites, err := siteClient.Get()
    if err != nil {
        log.Fatal(err)
    }

    for _, site := range sites {
        fmt.Printf("Site: %s (ID: %d)\n", site.Name, site.ID)
    }
}
```

## Authentication

The SDK uses cookie-based session authentication with two methods:

### Username/Password Authentication

```go
client, err := netris.Client(address, username, password, timeout)
if err != nil {
    log.Fatal(err)
}

// Authenticate and receive session cookie
err = client.Client.LoginUser()
if err != nil {
    log.Fatal(err)
}
```

### Session Cookie Authentication

```go
// Use existing session ID
client, err := netris.ClientWithCookie(address, sessionID, timeout)
if err != nil {
    log.Fatal(err)
}
```

### Development Mode (Insecure SSL)

```go
client, err := netris.Client(address, username, password, timeout)
client.Client.InsecureVerify(true) // Disable SSL verification (dev only!)
err = client.Client.LoginUser()
```

### Session Validation

```go
// Check if current session is valid
isValid, err := client.Client.CheckAuth()

// Start periodic auth checking (every 5 seconds)
go client.Client.CheckAuthWithInterval()
```

## Available Resources

### V2 API Resources (Recommended)

The V2 API provides enhanced functionality and more resources:

| Resource | Description | Client Method |
|----------|-------------|---------------|
| **Sites** | Network site management | `client.Site()` |
| **VPC** | Virtual Private Cloud management | `client.VPC()` |
| **VPC Peerings** | VPC peering connections | `client.VPCPeering()` |
| **VNet** | Virtual Network management | `client.VNet()` |
| **IPAM** | IP Address Management (allocations, subnets, hosts) | `client.IPAM()` |
| **L4LB** | Layer 4 Load Balancers | `client.L4LB()` |
| **NAT** | Network Address Translation | `client.NAT()` |
| **BGP** | Border Gateway Protocol configuration | `client.BGP()` |
| **Inventory** | Hardware inventory (switches, controllers, softgates, servers) | `client.Inventory()` |
| **Server Clusters** | Server cluster management | `client.ServerCluster()` |
| **Cluster Templates** | Server cluster templates | `client.ServerClusterTemplate()` |
| **Ports** | Switch port configuration | `client.Port()` |
| **Links** | Network link management | `client.Link()` |
| **DHCP** | DHCP option sets | `client.DHCP()` |
| **IP Reservations** | IP address reservations | `client.IPReservation()` |
| **VLAN Reservations** | VLAN reservations | `client.VLANReservation()` |
| **Public Key Ledger** | UFM Partition Keys Management | `client.PKEYLedger()` |
| **Version** | API version information | `client.Version()` |

### V1 API Resources

Legacy API with 21 resources including ACLs, users, permissions, tenants, and more. See [v1/api.go](v1/api.go) for full list.

## Usage Examples

### IPAM Management

```go
ipamClient := client.IPAM()

// List all IP allocations
allocations, err := ipamClient.Get()
if err != nil {
    log.Fatal(err)
}

// Get allocations for specific VPC
allocations, err := ipamClient.GetByVPC(vpcID)

// Get all subnets
subnets, err := ipamClient.GetSubnets()

// Get hosts in a subnet
hosts, err := ipamClient.GetHosts(subnetID)

// Create new allocation
import "github.com/netrisai/netriswebapi/v2/types/ipam"

allocation := &ipam.Allocation{
    Name:   "production-network",
    Prefix: "10.0.0.0/16",
    Tenant: ipam.IDName{ID: 1},
    Tags:   []string{"production", "datacenter-1"},
}

reply, err := ipamClient.AddAllocation(allocation)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Created allocation ID: %d\n", reply.Data.ID)

// Delete DHCP leases
reply, err := ipamClient.DeleteDHCPLease([]int{100, 101, 102})
```

### Site Management

```go
import "github.com/netrisai/netriswebapi/v2/types/site"

siteClient := client.Site()

// List all sites
sites, err := siteClient.Get()

// Get specific site details
siteDetail, err := siteClient.GetByID(1)

// Create new site
newSite := &site.Site{
    Name:         "datacenter-east",
    SwitchFabric: "netris",
    PublicASN:    65001,
    RohASN:       4200000000,
    VMASN:        4200000001,
    ACLDefaultPolicy: "permit",
}

reply, err := siteClient.Add(newSite)
if err != nil {
    log.Fatal(err)
}

// Update site
siteDetail.Name = "datacenter-east-updated"
reply, err = siteClient.Update(siteDetail.ID, siteDetail)

// Delete site
reply, err = siteClient.Delete(siteDetail.ID)
```

### VPC Operations

```go
import "github.com/netrisai/netriswebapi/v2/types/vpc"

vpcClient := client.VPC()

// List VPCs
vpcs, err := vpcClient.Get()

// Get VPC details
vpcDetail, err := vpcClient.GetByID(10)

// Create VPC
newVPC := &vpc.VPCw{
    Name: "production-vpc",
    Sites: []vpc.SiteGW{
        {
            ID:              1,
            GatewayIP:       "10.0.0.1/24",
            InventoryID:     5,
        },
    },
    Tenants: []int{1},
    Tags:    []string{"production"},
}

reply, err := vpcClient.Add(newVPC)

// Update VPC
vpcDetail.Name = "production-vpc-updated"
reply, err = vpcClient.Update(vpcDetail.ID, vpcDetail)

// Delete VPC
reply, err = vpcClient.Delete(vpcDetail.ID)
```

### Inventory Management

```go
import "github.com/netrisai/netriswebapi/v2/types/inventory"

inventoryClient := client.Inventory()

// List all hardware
hardware, err := inventoryClient.Get()

// Get specific device
device, err := inventoryClient.GetByID(5)

// Get installation command for agent
installCmd, err := inventoryClient.GetHwInstallAgent(5)
fmt.Println("Run this command on the server:", installCmd)

// Add a switch
newSwitch := &inventory.HWSwitchAdd{
    Name:        "leaf-switch-01",
    TenantID:    1,
    SiteID:      1,
    Description: "Production leaf switch",
    MainIP:      "192.168.1.10/24",
    MgmtIP:      "10.0.0.10/24",
    MacAddress:  "00:11:22:33:44:55",
    ProfileID:   1,
    PortCount:   48,
    Links: []inventory.HWLink{
        {
            Local:  "swp1",
            Remote: inventory.HWLinkRemote{
                ID:   2,
                Port: "swp1",
            },
        },
    },
}

reply, err := inventoryClient.AddSwitch(newSwitch)

// Delete hardware
reply, err = inventoryClient.Delete("switch", 5)
```

### Server Cluster Management

```go
import "github.com/netrisai/netriswebapi/v2/types/servercluster"

clusterClient := client.ServerCluster()

// List clusters
clusters, err := clusterClient.Get()

// Get cluster with resources (VNets, ROH, etc.)
cluster, err := clusterClient.GetByID(1)
fmt.Printf("Cluster: %s\n", cluster.Name)
fmt.Printf("VNets: %d\n", len(cluster.Resources.VNets))

// Create cluster from template
newCluster := &servercluster.ServerCluster{
    Name:                "k8s-production",
    SiteID:              1,
    ServerClusterID:     1, // template ID
    TenantID:            1,
    ServerClusterType:   "common",
    ServerClusterPolicy: "active-backup",
    Description:         "Production Kubernetes cluster",
}

reply, err := clusterClient.Add(newCluster)
```

### L4 Load Balancer

```go
import "github.com/netrisai/netriswebapi/v2/types/l4lb"

l4lbClient := client.L4LB()

// List load balancers
lbs, err := l4lbClient.Get()

// Create load balancer
newLB := &l4lb.L4LB{
    Name:        "web-lb",
    TenantID:    1,
    SiteID:      1,
    State:       "active",
    Protocol:    "tcp",
    IP:          "203.0.113.10",
    Port:        80,
    Check: l4lb.Check{
        Type:     "tcp",
        Timeout:  3000,
        Interval: 5000,
    },
    Backend: []l4lb.Backend{
        {
            IP:   "10.0.1.10",
            Port: 8080,
        },
        {
            IP:   "10.0.1.11",
            Port: 8080,
        },
    },
}

reply, err := l4lbClient.Add(newLB)
```

### BGP Configuration

```go
import "github.com/netrisai/netriswebapi/v2/types/bgp"

bgpClient := client.BGP()

// List BGP configurations
bgps, err := bgpClient.Get()

// Create BGP peering
newBGP := &bgp.BGP{
    Name:             "upstream-isp",
    SiteID:           1,
    InventoryID:      5,
    VNetID:           10,
    State:            "enabled",
    LocalIP:          "192.0.2.1",
    RemoteIP:         "192.0.2.2",
    LocalASN:         65001,
    RemoteASN:        65000,
    Description:      "Peering with ISP",
    Transport:        "default",
    BGPPassword:      "secret123",
    AllowAsIn:        0,
    DefaultOriginate: false,
    InboundRouteMap:  nil,
    OutboundRouteMap: nil,
    PrefixLimit: &bgp.PrefixLimit{
        MaxPrefixes: 100,
        Threshold:   80,
    },
}

reply, err := bgpClient.Add(newBGP)
```

## API Versions

### V2 API (Recommended)

```go
import netris "github.com/netrisai/netriswebapi/v2"

client, err := netris.Client(address, username, password, timeout)
```

**Benefits:**
- More resources (20 vs 21)
- Enhanced resource types with better structure
- Improved IPAM with subnet and host management
- VPC support
- Server cluster resources
- Better inventory management (separate types for switches, controllers, etc.)

### V1 API (Legacy)

```go
import netris "github.com/netrisai/netriswebapi/v1"

client, err := netris.Client(address, username, password, timeout)
```

## Error Handling

All API responses follow a consistent structure:

```go
reply, err := client.Site().Add(newSite)
if err != nil {
    // HTTP or network error
    log.Fatal(err)
}

// Check API-level success
if !reply.IsSuccess {
    fmt.Printf("API Error: %s\n", reply.Message)
    for _, err := range reply.Errors.Errors {
        fmt.Printf("  - %s\n", err)
    }
}

// Access returned data
siteID := reply.Data.ID
```

### Common Error Patterns

```go
// Network/HTTP errors
client, err := netris.Client(address, username, password, timeout)
if err != nil {
    log.Fatal("Failed to create client:", err)
}

// Authentication errors
err = client.Client.LoginUser()
if err != nil {
    log.Fatal("Failed to authenticate:", err)
}

// API operation errors
reply, err := siteClient.Add(newSite)
if err != nil {
    log.Fatal("Request failed:", err)
}

if !reply.IsSuccess {
    log.Printf("Operation failed: %s", reply.Message)
    if len(reply.Errors.Errors) > 0 {
        for _, e := range reply.Errors.Errors {
            log.Printf("  Error: %s", e)
        }
    }
}
```

## Best Practices

### 1. Use V2 API When Possible

```go
// Preferred
import netris "github.com/netrisai/netriswebapi/v2"
```

### 2. Set Appropriate Timeouts

```go
// For quick operations
client, err := netris.Client(address, user, pass, 10)

// For complex operations (large inventory, etc.)
client, err := netris.Client(address, user, pass, 60)
```

### 3. Secure SSL in Production

```go
// Development only
client.Client.InsecureVerify(true)

// Production - always verify SSL certificates
// (default behavior, no need to call InsecureVerify)
```

### 4. Implement Session Re-authentication

```go
// Check auth status periodically
isValid, err := client.Client.CheckAuth()
if err != nil || !isValid {
    err = client.Client.LoginUser()
    if err != nil {
        log.Fatal("Re-authentication failed:", err)
    }
}

// Or use automatic checking
go client.Client.CheckAuthWithInterval() // checks every 5 seconds
```

### 5. Handle Resources Properly

```go
// Always check before creating
sites, err := client.Site().Get()
// ... check if site exists ...

// Clean up on errors
reply, err := client.VPC().Add(newVPC)
if err != nil || !reply.IsSuccess {
    // Handle cleanup if needed
}
```

### 6. Use Structured Logging

```go
import "log/slog"

reply, err := client.Site().Add(newSite)
if err != nil {
    slog.Error("Failed to create site",
        "error", err,
        "site", newSite.Name,
    )
    return err
}

slog.Info("Site created successfully",
    "id", reply.Data.ID,
    "name", newSite.Name,
)
```

## Project Structure

```
netriswebapi/
├── http/                  # HTTP client implementation
│   ├── http.go            # Core HTTP client with auth
│   ├── types.go           # API response structures
│   └── addresses/         # API endpoint definitions
│       ├── v1/
│       └── v2/
├── v1/                   # V1 API client
│   ├── api.go            # Clientset and resource accessors
│   └── types/            # V1 resource type definitions
│       ├── site/
│       ├── vnet/
│       ├── bgp/
│       └── ...
├── v2/                   # V2 API client (recommended)
│   ├── api.go            # Clientset and resource accessors
│   └── types/            # V2 resource type definitions
│       ├── site/
│       ├── vpc/
│       ├── ipam/
│       ├── inventory/
│       └── ...
├── examples/             # Usage examples
└── README.md             # This file
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support

For issues, questions, or contributions, please visit the [GitHub repository](https://github.com/netrisai/netriswebapi).

## Additional Resources

- [Netris Documentation](https://www.netris.io/docs/)

