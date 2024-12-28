package yamlLoader

type Backend struct {
	URL    string `yaml:"url"`
	Weight int    `yaml:"weight"`
}
