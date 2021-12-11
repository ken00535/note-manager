package config

// Config is interface
type Config interface {
	GetDbOption() DbOption
	GetRdbAdress() string
	GetRdbPort() int
	GetRdbPassword() string
	GetUsername() []string
	GetPassword() []string
	GetSecret() string
}

// DbOption is db option
type DbOption struct {
	Address   string
	Port      int
	Username  string
	Password  string
	Mechanism string
}
