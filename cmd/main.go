package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yasastharinda9511/go_gateway_api/loadBalancer"
)

func main() {
	// Define the -c flag
	configFile := flag.String("c", "", "Path to the configuration file")
	flag.Parse()

	// Check if the config file flag was provided
	if *configFile == "" {
		fmt.Println("Usage: go run main.go -c <config-file>")
		os.Exit(1)
	}

	// Use the provided config file
	loadBalancerBuilder := loadBalancer.NewLoadBalancerBuilder(*configFile)
	loadBalancer, err := loadBalancerBuilder.Build()

	if err == nil {
		loadBalancer.Start()
	} else {
		fmt.Println("Error building load balancer:", err)
	}
}
