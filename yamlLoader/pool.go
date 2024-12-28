package yamlLoader

type Pool struct {
	LoadBalancer string    `yaml:"load_balancer"`
	Backends     []Backend `yaml:"backends"`
}
