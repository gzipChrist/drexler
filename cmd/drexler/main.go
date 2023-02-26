package main

import (
	"flag"
	"github.com/gzipchrist/drexler/internal/cli"
	"github.com/gzipchrist/drexler/internal/style"
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
