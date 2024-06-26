package main

import (
	"log"

	"github.com/camtrik/go-tools/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
