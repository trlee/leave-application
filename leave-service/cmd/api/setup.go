package main

import "log"

func setLog() {
	log.SetPrefix("[Leave-Service]: ")
	log.SetFlags(0)
}
