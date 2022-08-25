package main

import (
	"Crunch-Garage/go-video-call/internalFunctions/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
