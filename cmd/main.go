package main

import (
	"log"

	"github.com/yasastharinda9511/go_gateway_api/pipeline"
	"github.com/yasastharinda9511/go_gateway_api/pool"
	"github.com/yasastharinda9511/go_gateway_api/ruleStore"
	"github.com/yasastharinda9511/go_gateway_api/rules"
	"github.com/yasastharinda9511/go_gateway_api/server"
)

func main() {

	ruleStore := ruleStore.NewRuleStore()
	poolSelector := pool.NewPoolSelector()

	backend1 := pool.NewBackend("http://localhost:3000", 1)
	backend2 := pool.NewBackend("http://localhost:3001", 1)
	backend3 := pool.NewBackend("http://localhost:3002", 1)

	backends := []*pool.Backend{backend1, backend2, backend3}

	pool := pool.NewPool("abc", pool.WEIGHTEDLOADBALANCER, backends)
	poolSelector.AddPool(pool)

	hRule := rules.NewHeaderRule("Yasas", "tharinda")

	ruleStore.AddRule("abc", hRule)
	ruleStore.PrintAllRules()

	pipeline := pipeline.NewProcessingPipeline(ruleStore, poolSelector)
	srv := server.NewServer("3333", pipeline)

	// Register routes
	srv.RegisterRoutes()

	// Start the server
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
