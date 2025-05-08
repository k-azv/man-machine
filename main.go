package main

import (
	"log"
	"os"

	"github.com/k-azv/man-machine/config"
	flag "github.com/spf13/pflag"
)

func main() {
	var help, bare bool
	var iwant string
	flag.BoolVarP(&help, "help", "h", false, "Display help information")
	flag.StringVarP(&iwant, "iwant", "i", "", "Specify your needs for LLM to generate commands")
	flag.BoolVarP(&bare, "bare", "b", false, "Execute the provided command literally to fetch help documentation,\nbypassing mam's internal attempts")

	flag.Parse()

	commands := flag.Args() // command exclude flag and argument

	if help && !bare {
		printUsage()
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		log.Printf("Error: no command provided\n\n")
		printUsage()
		os.Exit(1)
	}

	if commands[0] == "setup" && !bare {
		runSetup()
		os.Exit(0)
	}

	var cfg config.Config
	var err error
	if cfg, err = config.Load(); err != nil {
		log.Fatalf("Error: loading config: %v", err)
	}

	pg := NewPromptGenerator(cfg)
	client := initClient(cfg)

	// Generate prompt depend on flag
	switch {
	case iwant != "":
		pg.GenerateIwant(iwant)
	default:
		pg.GenerateBasic()
	}

	var cmdDoc string
	// Generate document depend on flag
	if bare {
		var err error
		cmdDoc, err = bareFetchDoc(commands)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	} else {
		cmdDoc = fetchCmdDoc(commands)
	}

	if err := Chat(client, cmdDoc, pg, cfg); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

}
