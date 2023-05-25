package config

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const YAML_EXTENSION = ".yaml"

func ExportYAML(fp string, c *Config) error {
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(fp, out, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ExportJSON(fp string, c any) error {
	out, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(fp, out, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(fp string, c *Config) error {
	data, err := os.ReadFile(fp)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfigFiles(dp string) []Config {
	configs := []Config{}

	filepath.WalkDir(dp, func(s string, d fs.DirEntry, err error) error {
		if filepath.Ext(d.Name()) == YAML_EXTENSION {
			if err != nil {
				log.Println(err)
				return err
			}

			c := Config{}
			if err := LoadConfig(s, &c); err != nil {
				log.Println(err)
				return err
			}

			configs = append(configs, c)
		}

		return nil
	})

	return configs
}
