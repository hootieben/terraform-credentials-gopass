package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	gperr "github.com/gopasspw/gopass/pkg/action"
)

type resultJSON struct {
	Token string `json:"token"`
}

func getCred(c string) (string, error) {
	cmd := exec.Command("gopass", "show", "-f", "-o", c)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func formatOutput(c string) (jsonresult []byte, err error) {
	result := resultJSON{c}
	jsonresult, err = json.Marshal(result)

	return jsonresult, err
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: terraform-credentials-gopass get <hostname>")
		os.Exit(1)
	}

	switch args[0] {
	case "get":
		out, err := getCred(args[1])
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() == gperr.ExitDecrypt {
					log.Fatalf("Unable to find or decrypt secrect %s\n", args[1])
				}
			}
			log.Fatalf("Failed with %s\n", err)
		}
		result, err := formatOutput(out)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(result)
		os.Stdout.WriteString("\n")
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "The gopass credentials helper is not able to %s credentials.\n", args[0])
		os.Exit(1)
	}
}
