package env

import (
	envParser "github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type WorkingEnv struct {
	Testing    WorkingTestStructOne `env:"TESTING,file" envDefault:"../asset/test.json" envExpand:"true"`
	TestingTwo WorkingTestStructTwo `env:"TESTING_1,file" envDefault:"../asset/test_1.json" envExpand:"true"`
}

type WorkingTestStructOne struct {
	DummyInfo []TestSub `json:"dummyInfo" yaml:"dummyInfo"`
}

type WorkingTestSub struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

type WorkingTestStructTwo struct {
	DummyInfo []WorkingTestSub `json:"dummyInfo" yaml:"dummyInfo"`
}

// UnmarshalText - custom parser for parsing cloud-regions.json file
func (c *WorkingTestStructOne) UnmarshalText(data []byte) error {
	return yaml.Unmarshal(data, &c)
}

// UnmarshalText - custom parser for parsing cloud-regions.json file
func (c *WorkingTestStructTwo) UnmarshalText(data []byte) error {
	return yaml.Unmarshal(data, &c)
}

func WorkingNew() (*WorkingEnv, error) {
	var info WorkingEnv
	if err := envParser.Parse(&info); err != nil {
		return nil, errors.Wrap(err, "Unable to parse envs: ")
	}
	return &info, nil
}
