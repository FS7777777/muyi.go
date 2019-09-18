package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Enabled bool   `yaml:"enabled"` //yaml：yaml格式 enabled：属性的为enabled
	Path    string `yaml:"path"`
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
