package config

type (
	Configuration struct {
		Addr string
	}
)

func NewConfig() (Configuration, error) {
	// @todo load configs from properties
	conf := Configuration{
		Addr: ":8080",
	}
	return conf, nil
}
