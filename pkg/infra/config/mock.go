package config

type mockConfig struct {
	config
}

// InitMock mock config
func InitMock() Config {
	return &mockConfig{}
}
