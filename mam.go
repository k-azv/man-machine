package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/k-azv/man-machine/config"
	flag "github.com/spf13/pflag"
)

//go:embed templates/config.yaml
var configTmpl []byte

func runSetup() {
	cfgFile, err := config.GetConfigFilePath()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get config file path: %w", err))
		return
	}

	// Use template for config.yaml if config.yaml not exist in ~/.config/mam
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		if err := os.WriteFile(cfgFile, configTmpl, 0o600); err != nil {
			log.Fatalf("Write config.yaml template: %v", err)
		}
	}

	// Use default editor to open config.yaml
	editor := os.Getenv("EDITOR")
	if editor == "" {
		switch runtime.GOOS {
		case "windows":
			editor = "notepad"
		case "darwin":
			editor = "open"
		case "linux":
			editor = "xdg-open"
		default:
			log.Fatalf("We can't find editor for your OS\n" +
				"or you can edit ~/.config/mam/config.yaml manually")
		}
	}

	cmd := exec.Command(editor, cfgFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed open editor: %v\n"+
			"or you can edit ~/.config/mam/config.yaml manually", err)
	}

	fmt.Print("Setup complete!\n\n")
	printUsage()
}

func printUsage() {
	fmt.Print(
		`Usage: mam <command>
Use LLM to easily read command docs

Commands:
  setup -- Set up the configuration

Options:`, "\n")
	flag.PrintDefaults()
}

// fetchCmdDoc retrieves help documentation for a given command.
func fetchCmdDoc(command ...string) string {
	// Generate attempts to fetch help documentation(also suit for subcommands)
	var attempts [][]string

	man := []string{"man"}
	man = append(man, command...)
	attempts = append(attempts, man)

	help := command
	attempts = append(attempts, append(help, "--help"))
	attempts = append(attempts, append(help, "-h"))

	for _, args := range attempts {
		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		text := normalize(string(out))
		if err == nil && text != "" {
			return text
		}
	}

	fmt.Printf("No help documentation found for '%s' (no man page or --help output)\n"+
		"The command will be sent directly to LLM.\n\n", command)

	return strings.Join(command, " ")
}

// normalize line endings to \n
func normalize(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, "\r\n", "\n"))
}
