package yamlLoader

type Rules struct {
	ID          string       `yaml:"id"`
	HeaderRules []HeaderRule `yaml:"header_rules"`
	PathRule    PathRule     `yaml:"path_rule"`
	Pool        Pool         `yaml:"pool"`
	RewriteURL  RewriteURL   `yaml:"rewrite_url"`
}
