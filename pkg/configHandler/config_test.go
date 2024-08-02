package confighandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExtraData1 struct {
	Foo string `yaml:"foo" json:"foo"`
	Bar int    `yaml:"bar" json:"bar"`
}

type ExtraData struct {
	Email   string `yaml:"email" json:"email"`
	Counter int    `yaml:"counter" json:"counter"`
}

func Test_UnmarshalYAML(t *testing.T) {
	c, err := New[ExtraData1]("./testdata/config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "bar", c.ExtraData.Foo)
	assert.Equal(t, 1234, c.ExtraData.Bar)
}

func Test_UnmarshalJSON(t *testing.T) {
	c, err := New[ExtraData1]("./testdata/config.json")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "bar", c.ExtraData.Foo)
	assert.Equal(t, 1234, c.ExtraData.Bar)
}
