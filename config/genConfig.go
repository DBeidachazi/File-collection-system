package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	CreateDB CreateDB `yaml:"create_db"`
	FileSize FileSize `yaml:"file_size"`
}
