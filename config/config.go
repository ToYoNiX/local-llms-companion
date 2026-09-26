package config

import (
	"os"
	"fmt"
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
		fmt.Println("OLLAMA_URL is not set!")
		os.Exit(1)
	}

	cfg := Config {
		Ollama: OllamaConfig {
			Url: url,
		},
	}
	
	return cfg
}

