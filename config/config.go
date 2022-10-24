package config

type Config struct {
	ChannelSecret string
	ChannelToken  string
}

func NewConfig() *Config {
	return &Config{
		ChannelSecret: "",
		ChannelToken:  "",
	}
}
