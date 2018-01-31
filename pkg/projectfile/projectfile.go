package projectfile

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Project covers the top level project structure of our yaml
type Project struct {
	Name         string     `yaml:"name"`
	Owner        string     `yaml:"owner"`
	Version      string     `yaml:"version"`
	Environments string     `yaml:"environments"`
	Platforms    []Platform `yaml:"platforms"`
	Languages    []Language `yaml:"languages"`
	Variables    []Variable `yaml:"variables"`
	Hooks        []Hook     `yaml:"hooks"`
	Commands     []Command  `yaml:"commands"`
}

type Platform struct {
	Name         string  `yaml:"name"`
	Os           string  `yaml:"os"`
	Version      string  `yaml:"version"`
	Architecture string  `yaml:"architecture"`
	Params       []Param `yaml:"params"`
}

type Param struct {
	Libc            string `yaml:"libc"`
	CompilerName    string `yaml:"compiler-name"`
	CompilerVersion string `yaml:"compiler-version"`
	Debug           bool   `yaml:"debug"`
	BuildFlags      string `yaml:"build-flags"`
}

// Language covers the language structure, which goes under Project
type Language struct {
	Name        string     `yaml:"name"`
	Version     string     `yaml:"version"`
	Constraints Constraint `yaml:"constraints"`
	Params      []Param    `yaml:"params"`
	Packages    []Package  `yaml:"packages"`
}

// Constraint covers the constraint structure, which can go under almost any other struct
type Constraint struct {
	Platform    string `yaml:"platform"`
	Environment string `yaml:"environment"`
}

// Package covers the package structure, which goes under the language struct
type Package struct {
	Name        string     `yaml:"name"`
	Version     string     `yaml:"version"`
	Constraints Constraint `yaml:"constraints"`
	Params      []Param    `yaml:"params"`
}

// Variable covers the variable structure, which goes under Project
type Variable struct {
	Name        string     `yaml:"name"`
	Value       string     `yaml:"value"`
	Constraints Constraint `yaml:"constraints"`
}

// Hook covers the hook structure, which goes under Project
type Hook struct {
	Name        string     `yaml:"name"`
	Value       string     `yaml:"value"`
	Constraints Constraint `yaml:"constraints"`
}

// Command covers the command structure, which goes under Project
type Command struct {
	Name        string     `yaml:"name"`
	Value       string     `yaml:"value"`
	Constraints Constraint `yaml:"constraints"`
}

// Parse the given filepath, which should be the full path to an activestate.yaml file
func Parse(filepath string) (*Project, error) {
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	project := Project{}
	err = yaml.Unmarshal([]byte(dat), &project)

	return &project, err
}

// Write to the given filepath, which should be the full path to an activestate.yaml file
func Write(filepath string, project *Project) error {
	dat, err := yaml.Marshal(&project)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(dat))
	if err != nil {
		return err
	}

	return nil
}
