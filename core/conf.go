package core

import (
	"FengfengStudy/config"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")
	fmt.Println(c)
}
