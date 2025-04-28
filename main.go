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
	var iwant string
	flag.BoolVarP(&help, "help", "h", false, "Display help information")
	flag.StringVarP(&iwant, "iwant", "i", "", "Specify your needs for LLM to generate commands")
	flag.Parse()

	commands := flag.Args() // command exclude flag and argument

	if flag.NArg() == 0 && !help {
		log.Printf("Error: no command provided\n\n")
		printUsage()
		os.Exit(1)
	}

	switch {
	case iwant != "":
		prompt.GenerateIwant(iwant)
	case help:
		printUsage()
		os.Exit(0)
	case commands[0] == "setup":
		runSetup()
		os.Exit(0)
	default:
		prompt.GenerateBasic()
	}

	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error: loading config: %v", err)
	}

	client := initClient()

	cmdDoc := fetchCmdDoc(commands)
	if err := Chat(client, cmdDoc); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

}
