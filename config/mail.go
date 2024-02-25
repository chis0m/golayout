package config

import "go-layout/utils"

type MailConfig struct {
	Host       string `validate:"require"`
	Port       string `validate:"require"`
	Username   string
	Password   string
	Encryption string `validate:"require"`
	FromName   string `validate:"require"`
	FromEmail  string `validate:"require"`
}

func LoadMailConfig() (MailConfig, error) {
	return MailConfig{
		Host:       utils.Getenv("MAIL_HOST", "sandbox.smtp.mailtrap.io"),
		Port:       utils.Getenv("MAIL_PORT", "587"),
		Username:   utils.Getenv("MAIL_USERNAME", ""),
		Password:   utils.Getenv("MAIL_PASSWORD", ""),
		Encryption: utils.Getenv("MAIL_ENCRYPTION", "tls"),
		FromName:   utils.Getenv("MAIL_FROM_NAME", "GoLayout App"),
		FromEmail:  utils.Getenv("MAIL_FROM_ADDRESS", "app@example.com"),
	}, nil
}
