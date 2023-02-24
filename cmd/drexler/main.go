package main

import (
	"cli_apps/charm_apps/create-tui-app/internal/cli"
	"cli_apps/charm_apps/create-tui-app/internal/style"
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	t := time.Now()

	style.ShowHeader()

	flag.Parse()
	args := flag.Args()

	err := cli.HandleArgs(args)
	if err != nil {
		log.Fatalln(err)
	}

	style.ShowComplete(t)

	os.Exit(0)
}
