package main

type AmqpConfig struct {
	AmqpDSN            string `envconfig:"AMQP_DSN" required:"true"`
	QueueName          string `envconfig:"QUEUE_NAME" default:"ToNotificationService"`
	RegExchangeName    string `envconfig:"REG_EXCHANGE_NAME" default:"UserRegistrations"`
	NotifyExchangeName string `envconfig:"NOTIFY_EXCHANGE_NAME" default:"UserNotifications"`
}

type config struct {
	AmqpConfig
}
