package config

type ServerConfig struct {
	Address string
	Port    uint
}

// Orchestrator = 127.0.0.1 - 8010
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Address: "127.0.0.1",
		Port:    8010,
	}
}
