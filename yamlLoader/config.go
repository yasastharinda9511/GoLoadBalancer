package yamlLoader

type Config struct {
	Server Server `yaml:"server"`

	Rules []Rules `yaml:"rules"`
}
