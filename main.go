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

	commands := flag.Args() // command exclude flag and argument

	if len(commands) < 1 {
		printUsage()
		os.Exit(1)
	}

	switch {
	case help:
		printUsage()
		os.Exit(0)
	case commands[0] == "setup":
		runSetup()
		os.Exit(0)
	default:
		if err := config.LoadConfig(); err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		prompt.Initialize()
		client := initClient()
		cmdDoc := fetchCmdDoc(commands)

		if err := chat(client, cmdDoc); err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	}

}
