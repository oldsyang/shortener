package config

type UserConfig struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Local string `mapstructure:"local"`
}

type ServerConfig struct {
	Name           string     `mapstructure:"name"`
	Server         string     `mapstructure:"server"`
	UserConfig     UserConfig `mapstructure:"user_config"`
	DatabaseConfig Database   `mapstructure:"database"`
}

type Database struct {
	Type        string `mapstructure:"type"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	Name        string `mapstructure:"name"`
	MinConn     int    `mapstructure:"min_conn"`
	MaxConn     int    `mapstructure:"max_conn"`
	TablePrefix string `mapstructure:"table_prefix"`
}
