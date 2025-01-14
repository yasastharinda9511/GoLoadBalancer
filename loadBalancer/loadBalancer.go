package loadBalancer

import (
	"fmt"
	"log"
	"sync"

	"github.com/yasastharinda9511/go_gateway_api/message"
	"github.com/yasastharinda9511/go_gateway_api/pipeline"
	"github.com/yasastharinda9511/go_gateway_api/pool"
	"github.com/yasastharinda9511/go_gateway_api/ruleStore"
	"github.com/yasastharinda9511/go_gateway_api/rules"
	"github.com/yasastharinda9511/go_gateway_api/server"
	"github.com/yasastharinda9511/go_gateway_api/urlRewriter"
	"github.com/yasastharinda9511/go_gateway_api/yamlLoader"
)

type LoadBalancer struct {
	servers                     []*server.Server
	ruleStores                  []*ruleStore.RuleStore
	poolSelectors               []*pool.PoolSelector
	requestProcessingPipeline   []*pipeline.RequestProcessingPipeline
	responseProcessingPipelines []*pipeline.ResponseProcessingPipeline
	requestPools                []*message.Pool[*message.HttpRequestMessage]
	urlRewriters                []*urlRewriter.URLRewriter
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		servers:                     []*server.Server{},
		ruleStores:                  []*ruleStore.RuleStore{},
		poolSelectors:               []*pool.PoolSelector{},
		requestProcessingPipeline:   []*pipeline.RequestProcessingPipeline{},
		responseProcessingPipelines: []*pipeline.ResponseProcessingPipeline{},
		requestPools:                []*message.Pool[*message.HttpRequestMessage]{},
		urlRewriters:                []*urlRewriter.URLRewriter{},
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

func (lb *LoadBalancer) AddUrlRewriters(urlRewriter *urlRewriter.URLRewriter) {
	lb.urlRewriters = append(lb.urlRewriters, urlRewriter)
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

func (lb *LoadBalancer) GetURLRewriters() []*urlRewriter.URLRewriter {
	return lb.urlRewriters
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

	for i := 0; i < serverCount; i++ {

		ruleStore := ruleStore.NewRuleStore()
		poolSelector := pool.NewPoolSelector()
		urlRewriter := urlRewriter.NewURLRewriter()

		requestMessagePool := message.NewPool(func() *message.HttpRequestMessage {
			return message.NewHttpRequestMessage()
		})

		responseMessagePool := message.NewPool(func() *message.HttpResponseMessage {
			return message.NewHttpResponseMessage()
		})

		reponsePipeline := pipeline.NewResponseProcessingPipeline(requestMessagePool, responseMessagePool)
		reqpipeline := pipeline.NewRequestProcessingPipeline(ruleStore, poolSelector, reponsePipeline, requestMessagePool, responseMessagePool, urlRewriter)

		srv := server.NewServer(fmt.Sprint(basePort+i), reqpipeline, requestMessagePool)

		srv.RegisterRoutes()

		loadBalancer.AddServer(srv)
		loadBalancer.AddRuleStore(ruleStore)
		loadBalancer.AddPoolSelector(poolSelector)
		loadBalancer.AddRequestProcessingPipeline(reqpipeline)
		loadBalancer.AddResponseProcessingPipeline(reponsePipeline)
		loadBalancer.AddUrlRewriters(urlRewriter)
	}

	rules := cfg.Rules
	for _, rule := range rules {
		rule_id := rule.ID

		rewritePath := rule.RewriteURL.RewritePath

		for _, header := range rule.HeaderRules {
			b.addHeaderRule(loadBalancer.GetRuleStores(), rule_id, header.Key, header.Value)
		}

		pathRule := rule.PathRule

		b.addPathRule(loadBalancer.GetRuleStores(), rule_id, pathRule.Path, pathRule.Type)
		pool := rule.Pool
		loadBalancerType := pool.LoadBalancer
		backends := pool.Backends

		b.addPool(rule_id, loadBalancer.GetPoolSelectors(), loadBalancerType, backends)
		if rewritePath != "" {
			b.addURLRewrites(rule_id, rewritePath, loadBalancer.GetURLRewriters())
		}
	}

	return loadBalancer, nil
}

func (b *LoadBalancerBuilder) addHeaderRule(ruleStore []*ruleStore.RuleStore, ruleID string, key string, value string) {
	for _, rs := range ruleStore {

		headerRule := rules.NewHeaderRule(key, value)
		rs.AddRule(ruleID, headerRule)
	}
}

func (b *LoadBalancerBuilder) addPathRule(ruleStore []*ruleStore.RuleStore, ruleID string, path string, pathType string) {
	for _, rs := range ruleStore {

		newPath := path

		if pathType == "PREFIX" {
			newPath += "/*"

		}
		pathRule := rules.NewPathRule(newPath)
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
		fmt.Println("######################################################")
		fmt.Println("pool object backend count is ", len(pool.GetBackends()))
		fmt.Println("direct pool backend count is ", len(poolBackends))

	}
}

func (b *LoadBalancerBuilder) addURLRewrites(ruleId string, rewriteURL string, ruleWriters []*urlRewriter.URLRewriter) {
	for _, ur := range ruleWriters {
		ur.InsertRewriteURL(ruleId, rewriteURL)
	}
}
