package config

import (
	"os"
)

type RabbitMQConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       map[string]interface{}
}

func NewRabbitMQConfig() *RabbitMQConfig {
	return &RabbitMQConfig{
		Host:       os.Getenv("RABBITMQ_HOST"),
		Port:       os.Getenv("RABBITMQ_PORT"),
		User:       os.Getenv("RABBITMQ_USER"),
		Password:   os.Getenv("RABBITMQ_PASSWORD"),
		Name:       "hello",
		Durable:    false,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
}

type S3Config struct {
	Region          string
	Bucket          string
	AccessKeyID     string
	SecretAccessKey string
}

func NewS3Config() *S3Config {
	return &S3Config{
		Region:          os.Getenv("S3_REGION"),
		AccessKeyID:     os.Getenv("S3_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("S3_SECRET_ACCESS_KEY"),
	}
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
}

func NewPostgresConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}
}
