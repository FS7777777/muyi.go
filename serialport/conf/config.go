package conf

import (
	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

func ConfigInit() (config *YamlConfig, err error) {

	conf := new(YamlConfig)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	beego.BeeLogger.Debug("yamlFile:", yamlFile)
	if err != nil {
		beego.BeeLogger.Debug("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		beego.BeeLogger.Info("Unmarshal: %v", err)
	}
	beego.BeeLogger.Info("conf", conf)
	return conf, err
}
