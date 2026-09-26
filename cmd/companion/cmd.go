package main

import (
	"log"
	"github.com/ToYoNiX/local-llms-companion/config"
)

func main () {
	log.SetPrefix("companion -> ")
	
	// Env variables
	EnvCfg := config.LoadConfig()

	log.Println("Ollama url is:", EnvCfg.Ollama.Url)
}
