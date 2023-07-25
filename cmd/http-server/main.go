package main

import (
	"log"
	"os"
)


func main() {
	// set trial key for self-host users
	os.Setenv("ZIROOM_SECRET_KEY", "8xEMrWkBARcDDYQ")
	// init
	server, err := Initialize()
	if err != nil {
		log.Panic(err)
	}
	// TODO: gracefulStop
	server.Start()
}

	
	

