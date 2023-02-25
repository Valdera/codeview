package main

import (
	"codeview/config"
	"codeview/migration"
	"codeview/server"
	"log"

	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: need proc as argument")
	}

	cmd := os.Args[1]

	switch cmd {
	case "server":
		ApplicationStart()
	case "migrate-up-all":
		migration.MigrateUp()
	case "migrate-down-all":
		migration.MigrateDown()
	case "migrate-up":
		migration.MigrateStepUp()
	case "migrate-down":
		migration.MigrateStepDown()
	}
}

func ApplicationStart() {
	log.Println("Starting server...")

	cfg := config.Init()

	cfg.LoadFromEnv()

	server, err := server.Init(cfg)
	if err != nil {
		return
	}

	server.Start()
}
