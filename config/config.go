package config

import (
	"os"
	"log"
	// this automatically loads and Setenv the environment variables in the code
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Ollama OllamaConfig
}

type OllamaConfig struct {
	Url string
}

func LoadConfig () Config {
	url, exists := os.LookupEnv("OLLAMA_URL")
	if !exists {
		log.Fatal("OLLAMA_URL is not set!")
	}

	cfg := Config {
		Ollama: OllamaConfig {
			Url: url,
		},
	}
	
	return cfg
}

