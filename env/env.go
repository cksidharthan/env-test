package env

import (
	envParser "github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"reflect"
)

type Environment struct {
	Testing    TestStructOne `env:"TESTING" envDefault:"../asset/test.json" envExpand:"true"`
	TestingTwo TestStructTwo `env:"TESTING_1" envDefault:"../asset/test_1.json" envExpand:"true"`
}

type TestStructOne struct {
	DummyInfo []TestSub `json:"dummyInfo" yaml:"dummyInfo"`
}

type TestSub struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

type TestStructTwo struct {
	DummyInfo []TestSub `json:"dummyInfo" yaml:"dummyInfo"`
}

func New() (*Environment, error) {
	var info Environment
	if err := envParser.ParseWithFuncs(&info, map[reflect.Type]envParser.ParserFunc{
		reflect.TypeOf(TestStructOne{}): testStructOneParser,
		reflect.TypeOf(TestStructTwo{}): testStructTwoParser,
	}); err != nil {
		return nil, errors.Wrap(err, "Unable to parse envs: ")
	}
	return &info, nil
}

func readFile(name string) ([]byte, error) {
	s, err := filepath.Abs(name)
	if err != nil {
		return nil, err
	}
	d, err := os.ReadFile(s)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func testStructOneParser(fileName string) (interface{}, error) {
	d, err := readFile(fileName)
	if err != nil {
		return nil, err
	}
	var testStruct TestStructOne
	err = yaml.Unmarshal(d, &testStruct)
	if err != nil {
		return nil, err
	}
	return testStruct, nil
}

func testStructTwoParser(fileName string) (interface{}, error) {
	d, err := readFile(fileName)
	if err != nil {
		return nil, err
	}
	var testStruct TestStructTwo
	err = yaml.Unmarshal(d, &testStruct)
	if err != nil {
		return nil, err
	}
	return testStruct, nil
}
