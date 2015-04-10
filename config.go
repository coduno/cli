package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Prepare []string `yaml:"prepare,flow"`
	Build   []string `yaml:"build,flow"`
	Run     string   `yaml:"run"`
}

var fileNames = [2]string{"coduno.yaml", "coduno.yml"}

func parseConfiguration() (*Config, error) {
	lim := "/"
	wd, err := os.Getwd()

	if err != nil {
		wd = "."
	}

	vol := filepath.VolumeName(wd)
	if vol != "" {
		lim = vol
	}
	return findConfigurationRecursive(wd, lim)
}

func findConfigurationRecursive(dirName, limit string) (*Config, error) {
	// In the current directory, try to open the config files
	for i := range fileNames {
		rawConfig, err := ioutil.ReadFile(path.Join(dirName, fileNames[i]))
		if err != nil {
			// TODO(flowlo): This might be interesting to log later
			//fmt.Printf(err.Error()+"\n")
			continue
		}
		var config Config
		err = yaml.Unmarshal(rawConfig, &config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing %s: %s\n", fileNames[i], err.Error())
			continue
		}
		return &config, nil
	}

	abs, err := filepath.Abs(dirName)
	if err != nil {
		return nil, err
	}
	if abs == limit {
		return nil, fmt.Errorf("reached %s before finding a valid configuration file.", limit)
	}

	return findConfigurationRecursive(path.Join(dirName, ".."), limit)
}
