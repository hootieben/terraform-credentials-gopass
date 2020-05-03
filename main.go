package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type resultJSON struct {
	Token string `json:"token"`
}

func valArgs() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println()
	}
}

func getCred(c string) (string, error) {
	cmd := exec.Command("gopass", c)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func formatOutput(c string) (resultJSON []byte, err error) {
	result := resultJSON{c}
	resultJSON, err = json.Marshal(result)
	return resultJSON, err
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: terraform-credentials-gopass get <hostname>")
		os.Exit(1)
	}
	switch args[0] {
	case "get":
		//cmd := exec.Command("gopass", args[1])
		//out, err := cmd.CombinedOutput()
		out, err := getCred(args[1])
		if err != nil {
			log.Fatalf("Failed with %s\n", err)
		}
		//result := resultJSON{string(out)}
		//resultJSON, err := json.Marshal(result)
		result, err := formatOutput(out)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(result)
		os.Stdout.WriteString("\n")
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "The 'env' credentials helper is not able to %s credentials.\n", args[0])
		os.Exit(1)
	}
}
