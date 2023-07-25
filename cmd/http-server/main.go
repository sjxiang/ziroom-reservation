package main

import "log"


func main() {
	// init
	server, err := Initialize()
	if err != nil {
		log.Panic(err)
	}
	// TODO: gracefulStop
	server.Start()
}

	
	

