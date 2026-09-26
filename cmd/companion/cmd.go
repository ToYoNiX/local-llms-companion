package main

import (
	"os"
	"log"
)

func main () {
	log.SetPrefix("companion -> ")

	url, exist := os.LookupEnv("OLLAMA_URL")

	if !exist {
		log.Fatal("OLLAMA_URL env variable is not set!")
	}

	log.Println("Ollama url is:", url)

	
}
