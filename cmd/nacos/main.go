package main

import (
	"log"
	"os"
	"registry-test/pkg/nacos"
)

func main() {
	if err := nacos.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
