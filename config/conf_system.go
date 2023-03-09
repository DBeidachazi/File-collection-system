package config

import "fmt"

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s *System) GetHost() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
