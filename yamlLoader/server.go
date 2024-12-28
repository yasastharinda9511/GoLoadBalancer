package yamlLoader

type Server struct {
	BasePort    int `yaml:"base_port"`
	ServerCount int `yaml:"server_count"`
}
