package main

import (
	"log"
	"os"
	"time"
)

const (
	APP_NAME = "EM-server"
)

func main() {
	logger := log.New(os.Stdout, "["+APP_NAME+"]", log.Ldate|log.Ltime)
	logger.Print("Starting...")
	time.Sleep(time.Second * 5)
	logger.Print("Done...")
}
