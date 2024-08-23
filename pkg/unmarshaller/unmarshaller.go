package unmarshaller

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type Extension string

const (
	JSON Extension = ".json"
	YAML Extension = ".yaml"
	YML  Extension = ".yml"
	TOML Extension = ".toml"
)

type Unmarshaller interface {
	Unmarshal(config interface{}) error
}

type JsonUnmarshaller struct {
	Data []byte
}

type TomlUnmarshaller struct {
	Data []byte
}

type YamlUnmarshaller struct {
	Data []byte
}

func (y *YamlUnmarshaller) Unmarshal(config interface{}) error {
	if err := yaml.Unmarshal(y.Data, config); err != nil {
		return err
	}
	return nil
}

func (j *JsonUnmarshaller) Unmarshal(config interface{}) error {
	if err := json.Unmarshal(j.Data, config); err != nil {
		return err
	}
	return nil
}

func (t *TomlUnmarshaller) Unmarshal(config interface{}) error {
	if err := toml.Unmarshal(t.Data, config); err != nil {
		return err
	}
	return nil
}

// NewUnmarshaller FactoryPattern function to create the appropriate Unmarshaller based on the file extension
func NewUnmarshaller(path string) (Unmarshaller, error) {
	ext := filepath.Ext(path)
	payload, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	switch Extension(ext) {
	case JSON:
		return &JsonUnmarshaller{
			Data: payload,
		}, nil
	case TOML:
		return &TomlUnmarshaller{
			Data: payload,
		}, nil
	case YAML, YML:
		return &YamlUnmarshaller{
			Data: payload,
		}, nil
	default:
		return nil, errors.New("unsupported file extension")
	}
}
