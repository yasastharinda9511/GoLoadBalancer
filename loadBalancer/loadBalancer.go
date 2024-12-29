package loadBalancer

import (
	"fmt"
	"log"
	"sync"

	"github.com/yasastharinda9511/go_gateway_api/pipeline"
	"github.com/yasastharinda9511/go_gateway_api/pool"
	"github.com/yasastharinda9511/go_gateway_api/ruleStore"
	"github.com/yasastharinda9511/go_gateway_api/rules"
	"github.com/yasastharinda9511/go_gateway_api/server"
	"github.com/yasastharinda9511/go_gateway_api/yamlLoader"
)

type LoadBalancer struct {
	servers                     []*server.Server
	ruleStores                  []*ruleStore.RuleStore
	poolSelectors               []*pool.PoolSelector
	requestProcessingPipeline   []*pipeline.RequestProcessingPipeline
	responseProcessingPipelines []*pipeline.ResponseProcessingPipeline
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		servers:                     []*server.Server{},
		ruleStores:                  []*ruleStore.RuleStore{},
		poolSelectors:               []*pool.PoolSelector{},
		requestProcessingPipeline:   []*pipeline.RequestProcessingPipeline{},
		responseProcessingPipelines: []*pipeline.ResponseProcessingPipeline{},
	}
}

func (lb *LoadBalancer) AddServer(server *server.Server) {
	lb.servers = append(lb.servers, server)
}

func (lb *LoadBalancer) AddRuleStore(ruleStore *ruleStore.RuleStore) {
	lb.ruleStores = append(lb.ruleStores, ruleStore)
}

func (lb *LoadBalancer) AddPoolSelector(poolSelector *pool.PoolSelector) {
	lb.poolSelectors = append(lb.poolSelectors, poolSelector)
}

func (lb *LoadBalancer) AddRequestProcessingPipeline(requestProcessingPipeline *pipeline.RequestProcessingPipeline) {
	lb.requestProcessingPipeline = append(lb.requestProcessingPipeline, requestProcessingPipeline)
}

func (lb *LoadBalancer) AddResponseProcessingPipeline(responseProcessingPipeline *pipeline.ResponseProcessingPipeline) {
	lb.responseProcessingPipelines = append(lb.responseProcessingPipelines, responseProcessingPipeline)
}

func (lb *LoadBalancer) GetServers() []*server.Server {
	return lb.servers
}

func (lb *LoadBalancer) GetRuleStores() []*ruleStore.RuleStore {
	return lb.ruleStores
}

func (lb *LoadBalancer) GetPoolSelectors() []*pool.PoolSelector {
	return lb.poolSelectors
}

func (lb *LoadBalancer) GetRequestProcessingPipelines() []*pipeline.RequestProcessingPipeline {
	return lb.requestProcessingPipeline
}

func (lb *LoadBalancer) GetResponseProcessingPipelines() []*pipeline.ResponseProcessingPipeline {
	return lb.responseProcessingPipelines
}

func (lb *LoadBalancer) Start() {
	var wg sync.WaitGroup
	for _, srv := range lb.servers {
		wg.Add(1)
		go func(s *server.Server) {
			defer wg.Done()
			if err := s.Start(); err != nil {
				log.Fatalf("Failed to start server: %v", err)
			}
		}(srv)
	}
	wg.Wait()
}

type LoadBalancerBuilder struct {
	configFile string
}

func NewLoadBalancerBuilder(configFile string) *LoadBalancerBuilder {
	return &LoadBalancerBuilder{configFile: configFile}
}

func (b *LoadBalancerBuilder) Build() (*LoadBalancer, error) {
	fmt.Println("LoadBalancer Build Called !!!")
	yamlLoader := yamlLoader.NewYamlLoader()
	// Load the configuration
	cfg, err := yamlLoader.LoadConfig(b.configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	loadBalancer := NewLoadBalancer()

	basePort := cfg.Server.BasePort
	serverCount := cfg.Server.ServerCount
	fmt.Println(serverCount)

	for i := 0; i < serverCount; i++ {

		ruleStore := ruleStore.NewRuleStore()
		poolSelector := pool.NewPoolSelector()
		reponsePipeline := pipeline.NewResponseProcessingPipeline()
		reqpipeline := pipeline.NewRequestProcessingPipeline(ruleStore, poolSelector, reponsePipeline)
		srv := server.NewServer(fmt.Sprint(basePort+i), reqpipeline)

		srv.RegisterRoutes()

		loadBalancer.AddServer(srv)
		loadBalancer.AddRuleStore(ruleStore)
		loadBalancer.AddPoolSelector(poolSelector)
		loadBalancer.AddRequestProcessingPipeline(reqpipeline)
		loadBalancer.AddResponseProcessingPipeline(reponsePipeline)
	}

	fmt.Println("Rules Count : ", len(cfg.Rules))
	rules := cfg.Rules
	fmt.Print(len(rules))
	for _, rule := range rules {
		rule_id := rule.ID

		for _, header := range rule.HeaderRules {
			fmt.Print(len(loadBalancer.GetRuleStores()))
			b.addHeaderRule(loadBalancer.GetRuleStores(), rule_id, header.Key, header.Value)

		}

		pathRule := rule.PathRule

		b.addPathRule(loadBalancer.GetRuleStores(), rule_id, pathRule.Path, pathRule.Type)
		pool := rule.Pool
		loadBalancerType := pool.LoadBalancer
		backends := pool.Backends

		b.addPool(rule_id, loadBalancer.GetPoolSelectors(), loadBalancerType, backends)

	}

	return loadBalancer, nil
}

func (b *LoadBalancerBuilder) addHeaderRule(ruleStore []*ruleStore.RuleStore, ruleID string, key string, value string) {
	for _, rs := range ruleStore {

		headerRule := rules.NewHeaderRule(key, value)
		headerRule.Print()
		rs.AddRule(ruleID, headerRule)
	}
}

func (b *LoadBalancerBuilder) addPathRule(ruleStore []*ruleStore.RuleStore, ruleID string, path string, pathType string) {
	for _, rs := range ruleStore {

		if pathType == "prefix" {
			path += "/*"
		}
		pathRule := rules.NewPathRule(path)
		pathRule.Print()
		rs.AddRule(ruleID, pathRule)
	}
}

func (b *LoadBalancerBuilder) addPool(ruleId string, poolSelector []*pool.PoolSelector, loadbalancertype string, backends []yamlLoader.Backend) {

	for _, ps := range poolSelector {

		lbType, err := pool.ParseLoadBalancerType(loadbalancertype)
		if err != nil {
			log.Fatalf("Failed to parse load balancer type: %v", err.Error())
			return
		}

		poolBackends := []*pool.Backend{}
		for _, backend := range backends {
			poolBackends = append(poolBackends, pool.NewBackend(backend.URL, backend.Weight))
		}

		pool := pool.NewPool(ruleId, lbType, poolBackends)
		ps.AddPool(pool)
	}
}
