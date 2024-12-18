package config

type Config struct {
	DB    DBConfig    `json:"db"`
	Bot   BotConfig   `json:"bot"`
	Redis RedisConfig `json:"redis"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type BotConfig struct {
	Token string `json:"token"`
	AppID string `json:"app_id"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
