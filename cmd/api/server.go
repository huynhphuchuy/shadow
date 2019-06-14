package main

import (
	"flag"
	"fmt"
	"os"
	"server/cmd/api/routes"
	"server/internal/config"
	"server/internal/platform/mongo"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	mongo.Init()
	routes.Init()
}
