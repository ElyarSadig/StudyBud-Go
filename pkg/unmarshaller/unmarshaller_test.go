package unmarshaller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyConfig struct {
	FooBar string `yaml:"foo_bar" json:"foo_bar"`
	Bar    string `yaml:"bar" json:"bar"`
}

type Company struct {
	Name  string
	Phone string
}
type Person struct {
	Name    string
	Age     int64
	Company Company
}

func TestYAMLUnmarshaller_Unmarshal(t *testing.T) {
	yamlData := []byte(`
foo_bar: test
bar: 123
`)

	config := &MyConfig{}
	unmarshaller := &yamlUnmarshaller{Data: yamlData}

	err := unmarshaller.Unmarshal(config)
	if err != nil {
		t.Errorf("Failed to unmarshal YAML: %s", err)
	}

	expectedConfig := &MyConfig{
		FooBar: "test",
		Bar:    "123",
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestTomlUnmarshaller_Unmarshal(t *testing.T) {
	tomlData := []byte(`
	name = "John Doe"
	age = 42
	[company]
		name = "Company"
		phone = "+1 9123456789"
	`)

	config := &Person{}
	unmarshaller := &tomlUnmarshaller{Data: tomlData}

	err := unmarshaller.Unmarshal(config)
	if err != nil {
		t.Errorf("Failed to unmarshal YAML: %s", err)
	}

	expectedConfig := &Person{
		Name: "John Doe",
		Age:  42,
		Company: Company{
			Name:  "Company",
			Phone: "+1 9123456789",
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestJSONUnmarshaller_Unmarshal(t *testing.T) {
	jsonData := []byte(`
		{ "foo_bar": "test","bar":"123" }
	`)

	expectedConfig := &MyConfig{
		FooBar: "test",
		Bar:    "123",
	}

	config := &MyConfig{}
	unmarshaller := &jsonUnmarshaller{Data: jsonData}

	err := unmarshaller.Unmarshal(config)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %s", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestCreateUnmarshaller(t *testing.T) {
	t.Run("JSON extension", func(t *testing.T) {
		path := "./testdata/test.json"

		unmarshaller, err := NewUnmarshaller(path)
		assert.NoError(t, err)
		assert.IsType(t, &jsonUnmarshaller{}, unmarshaller)
	})

	t.Run("TOML extension", func(t *testing.T) {
		path := "./testdata/test.toml"

		unmarshaller, err := NewUnmarshaller(path)
		assert.NoError(t, err)
		assert.IsType(t, &tomlUnmarshaller{}, unmarshaller)
	})

	t.Run("YAML extension", func(t *testing.T) {
		path := "./testdata/test.yaml"

		unmarshaller, err := NewUnmarshaller(path)
		assert.NoError(t, err)
		assert.IsType(t, &yamlUnmarshaller{}, unmarshaller)
	})

	t.Run("Unsupported extension", func(t *testing.T) {
		path := "/path/to/file.txt"

		unmarshaller, err := NewUnmarshaller(path)
		assert.Error(t, err)
		assert.Nil(t, unmarshaller)
	})
}
