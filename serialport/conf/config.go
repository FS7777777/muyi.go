package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type YamlConfig struct {
	Serial struct {
		Name    string `yaml:"name"`
		Baud    int    `yaml:"baud"`
		Timeout int    `yaml:"timeout"`
		Command []byte `yaml:"command"`
	}
	Rabbit struct {
		Amqp     string `yaml:"amqp"`
		Exchange string `yaml:"exchange"`
	}
}

func ConfigInit() (config *YamlConfig) {

	conf := new(YamlConfig)
	yamlFile, err := ioutil.ReadFile("config.yaml")

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
	return conf
}
