package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type portConfig struct {
	Enabled        bool   `yaml:"enabled"`   //yaml：yaml格式 enabled：属性的为enabled
	HttpPort       string `yaml:"http_port"` //yaml：yaml格式 enabled：属性的为enabled
	TMPort         string `yaml:"tm_port"`
	ImagePort      string `yaml:"image_port"`
	TCLoopbackPort string `yaml:"tc_loopback_port"`
	TCPort         string `yaml:"tc_port"`
	VoicePort      string `yaml:"voice_port"`
}

var (
	PortConfig = (&portConfig{}).getConf()
)

func (c *portConfig) getConf() *portConfig {
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
