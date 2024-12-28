package yamlLoader

type Rules struct {
	ID          string       `yaml:"id"`
	HeaderRules []HeaderRule `yaml:"header_rules"`
	Pool        Pool         `yaml:"pool"`
}
