package config

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`

	JwtSecret struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwtSecret"`
}
