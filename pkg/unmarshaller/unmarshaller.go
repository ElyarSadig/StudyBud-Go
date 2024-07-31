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

type jsonUnmarshaller struct {
	Data []byte
}

type tomlUnmarshaller struct {
	Data []byte
}

type yamlUnmarshaller struct {
	Data []byte
}

func (y *yamlUnmarshaller) Unmarshal(config interface{}) error {
	if err := yaml.Unmarshal(y.Data, config); err != nil {
		return err
	}
	return nil
}

func (j *jsonUnmarshaller) Unmarshal(config interface{}) error {
	if err := json.Unmarshal(j.Data, config); err != nil {
		return err
	}
	return nil
}

func (t *tomlUnmarshaller) Unmarshal(config interface{}) error {
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
		return &jsonUnmarshaller{
			Data: payload,
		}, nil
	case TOML:
		return &tomlUnmarshaller{
			Data: payload,
		}, nil
	case YAML, YML:
		return &yamlUnmarshaller{
			Data: payload,
		}, nil
	default:
		return nil, errors.New("unsupported file extension")
	}
}
