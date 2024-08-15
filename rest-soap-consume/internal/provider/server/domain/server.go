package domain

type ServerConfig struct {
	Address string
	Port    uint
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Address: "127.0.0.1",
		Port:    8080,
	}
}
