package main

type DbConfig struct {
	DbDriver string `envconfig:"DB_DRIVER" default:"postgres"`
	DbDSN    string `envconfig:"DB_DSN" required:"true"`
}

type AmqpConfig struct {
	AmqpDSN         string `envconfig:"AMQP_DSN" required:"true"`
	RegExchangeName string `envconfig:"REG_EXCHANGE_NAME" default:"UserRegistrations"`
}

type config struct {
	DbConfig
	AmqpConfig
	ServerAddr string `envconfig:"SERVER_ADDR" required:"true"`
}
