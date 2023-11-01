package config

type Swagger struct {
	Title    string `mapstructure:"title" json:"title" yaml:"title"`
	Desc     string `mapstructure:"desc" json:"desc" yaml:"desc"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	BasePath string `mapstructure:"base_path" json:"base_path" yaml:"base_path"`
}
