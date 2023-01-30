package config

type DB struct {
	Driver     string `yaml:"Driver"`
	DataSource string `yaml:"DataSource"`
}

type Redis struct {
	Host string `yaml:"Host"`
	Pass string `yaml:"Pass"`
}

type Log struct {
	StorageLocation     string `yaml:"StorageLocation"`
	RotationTime        int    `yaml:"RotationTime"`
	RemainRotationCount uint   `yaml:"RemainRotationCount"`
}

type Config struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
	Mode string `yaml:"Mode"`

	DB DB `yaml:"DB"`

	Redis Redis `yaml:"Redis"`

	Log Log `yaml:"Log"`
}
