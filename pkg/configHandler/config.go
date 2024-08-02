package confighandler

import "github.com/elyarsadig/studybud-go/pkg/unmarshaller"

func New[Extra any](configPath string) (*Config[Extra], error) {
	return newUnmarshal[Extra](configPath)
}

func newUnmarshal[Extra any](configPath string) (*Config[Extra], error) {
	cfg := new(Config[Extra])

	unmarshal, err := unmarshaller.NewUnmarshaller(configPath)
	if err != nil {
		return nil, err
	}

	return cfg, unmarshal.Unmarshal(cfg)
}
