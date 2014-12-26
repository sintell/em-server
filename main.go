package main

import (
	"github.com/sintell/em-server/server"
	"log"
	"os"
)

const (
	APP_NAME = "EM-server"
)

func main() {
	logger := log.New(os.Stdout, "["+APP_NAME+"]", log.Ldate|log.Ltime)
	logger.Print("Starting...")

	stop := make(chan bool)

	srv := server.New(":8090", server.ERRORS)
	srv.Start()

	logger.Print("Done...")
	<-stop
}
