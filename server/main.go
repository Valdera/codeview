package main

import (
	"codeview/migration"
	"codeview/server"

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
		server.Start()
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
