package yamlLoader

import (
	"os"

	"gopkg.in/yaml.v2"
)

type YamlLoader struct{}

func NewYamlLoader() *YamlLoader {
	return &YamlLoader{}
}

// LoadConfig reads a YAML configuration file using the os package and parses it into the Config struct.
func (y *YamlLoader) LoadConfig(filePath string) (*Config, error) {
	// Open the YAML file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get file info to determine the size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Read the file into a byte slice
	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	// Parse the YAML into the Config struct
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
