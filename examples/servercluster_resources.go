package main

import (
	"fmt"
	"log"

	napi "github.com/netrisai/netriswebapi/v2"
	_ "github.com/netrisai/netriswebapi/v2/types/servercluster"
)

func main() {
	// Configuration
	var (
		url      = "http://dev.netris.dev/"
		username = "username"
		password = "password"
		timeout  = 100
	)

	// Create Netris client
	client, err := napi.Client(url, username, password, timeout)
	if err != nil {
		log.Fatalf("Failed to create Netris client: %v", err)
	}

	// Enable insecure SSL verification for development
	client.Client.InsecureVerify(true)

	// Login to Netris
	fmt.Println("Logging in to Netris...")
	err = client.Client.LoginUser()
	if err != nil {
		log.Fatalf("Failed to login to Netris: %v", err)
	}
	fmt.Println("Successfully logged in to Netris")

	// Get ServerCluster client
	serverClusterClient := client.ServerCluster()

	// Example 1: Get all server clusters
	fmt.Println("\n=== Getting all server clusters ===")
	clusters, err := serverClusterClient.Get()
	if err != nil {
		log.Fatalf("Failed to get server clusters: %v", err)
	}

	fmt.Printf("Found %d server clusters:\n", len(clusters))
	for _, cluster := range clusters {
		fmt.Printf("  - ID: %d, Name: %s\n", cluster.ID, cluster.Name)
	}

	if len(clusters) == 0 {
		fmt.Println("No server clusters found. Please create one first.")
		return
	}

	// Example 2: Get server cluster by ID with resources
	fmt.Printf("\n=== Getting server cluster by ID with resources ===\n")
	clusterID := clusters[0].ID // Use the first cluster
	fmt.Printf("Getting server cluster ID: %d\n", clusterID)

	cluster, err := serverClusterClient.GetByID(clusterID)
	if err != nil {
		log.Fatalf("Failed to get server cluster by ID: %v", err)
	}

	fmt.Printf("Server Cluster Details:\n")
	fmt.Printf("  ID: %d\n", cluster.ID)
	fmt.Printf("  Name: %s\n", cluster.Name)
	fmt.Printf("  State: %s\n", cluster.State)
	fmt.Printf("  Status: %s (%s)\n", cluster.Status.Label, cluster.Status.Value)
	fmt.Printf("  Admin: %s (ID: %d)\n", cluster.Admin.Name, cluster.Admin.ID)
	fmt.Printf("  Site: %s (ID: %d)\n", cluster.Site.Name, cluster.Site.ID)
	fmt.Printf("  VPC: %s (ID: %d)\n", cluster.VPC.Name, cluster.VPC.ID)
	fmt.Printf("  Template: %s (ID: %d)\n", cluster.SrvClusterTemplate.Name, cluster.SrvClusterTemplate.ID)
	fmt.Printf("  Tags: %v\n", cluster.Tags)
	fmt.Printf("  Servers: %d servers\n", len(cluster.Servers))

	// Display servers
	for i, server := range cluster.Servers {
		fmt.Printf("    Server %d: %s (ID: %d, Shared: %t)\n", i+1, server.Name, server.ID, server.Shared)
	}

	// Example 3: Access the structured resources field
	fmt.Printf("\n=== Server Cluster Resources ===\n")
	fmt.Printf("VNets: %d found\n", len(cluster.Resources.VNets))
	for i, vnet := range cluster.Resources.VNets {
		fmt.Printf("  VNet %d: %s (ID: %d)\n", i+1, vnet.Name, vnet.ID)
		fmt.Printf("    IPv4 Gateways: %d found\n", len(vnet.IPv4Gateways))
		for j, gateway := range vnet.IPv4Gateways {
			fmt.Printf("      Gateway %d: %s\n", j+1, gateway.Prefix)
		}
		fmt.Printf("    IPv6 Gateways: %d found\n", len(vnet.IPv6Gateways))
		for j, gateway := range vnet.IPv6Gateways {
			fmt.Printf("      Gateway %d: %s\n", j+1, gateway.Prefix)
		}
	}

	fmt.Printf("\nAllocations: %d found\n", len(cluster.Resources.Allocations))
	for i, allocation := range cluster.Resources.Allocations {
		fmt.Printf("  Allocation %d: %s (ID: %d)\n", i+1, allocation.Prefix, allocation.ID)
	}

	fmt.Printf("\nSubnets: %d found\n", len(cluster.Resources.Subnets))
	for i, subnet := range cluster.Resources.Subnets {
		fmt.Printf("  Subnet %d: %s (ID: %d)\n", i+1, subnet.Prefix, subnet.ID)
	}

	// Example 4: Display template gateway information
	fmt.Printf("\n=== Server Cluster Template Gateways ===\n")
	fmt.Printf("Template: %s\n", cluster.SrvClusterTemplate.Name)

	fmt.Println("\n=== Example completed successfully! ===")
	fmt.Println("The ServerCluster.GetByID() method now includes the .resources map in its response.")
}
