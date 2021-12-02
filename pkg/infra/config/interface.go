package config

// Config is interface
type Config interface {
	GetDbAddress() string
	GetDbPort() int
	GetRdbAdress() string
	GetRdbPort() int
	GetRdbPassword() string
	GetUsername() []string
	GetPassword() []string
	GetSecret() string
}
