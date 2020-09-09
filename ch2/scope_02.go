package main

import "os"
import "log"

var cwd string

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
