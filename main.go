package main

import (
	"log"
	"os"

	"github.com/k-azv/man-machine/config"
	"github.com/k-azv/man-machine/prompt"
	flag "github.com/spf13/pflag"
)

func main() {
	var help bool
	flag.BoolVarP(&help, "help", "h", false, "Display help information")
	flag.Parse()

	switch {
	case help:
		printUsage()
	case len(os.Args) > 1 && os.Args[1] == "setup":
		runSetup()
	default:
		if err := config.LoadConfig(); err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		if len(os.Args) < 2 {
			printUsage()
			os.Exit(1)
		}

		prompt.Initialize()
		client := initClient()
		helpText := fetchHelp(os.Args[1])

		if err := chat(client, helpText); err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	}

}
