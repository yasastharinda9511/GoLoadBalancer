package main

import (
	"github.com/yasastharinda9511/go_gateway_api/loadBalancer"
)

func main() {

	loadBalancerBuilder := loadBalancer.NewLoadBalancerBuilder("config.yaml")
	loadBalancer, err := loadBalancerBuilder.Build()

	if err == nil {
		loadBalancer.Start()
	}
}
