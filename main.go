package main

import (
	"./cmd"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("DVSTRG> ")
	log.SetOutput(os.Stderr)
}

func main() {
	cmd.Execute()
}
