package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var config Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	// enitialize env loading
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// version
	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("version is required: ", err)
		os.Exit(1)
	}

	// serviceName
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("service Name is required: ", err)
		os.Exit(1)
	}

	// HTTP_PORT
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("HTTP PORT service Name is required: ", err)
		os.Exit(1)
	}

	httpInt, err := strconv.Atoi(httpPort)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
		os.Exit(1)
	}

	config = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpInt,
	}

}

func GetConfig() *Config {
	loadConfig()
	return &config
}
