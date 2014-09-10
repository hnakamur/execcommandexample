package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/mgutz/ansi"
)

func main() {
	cmd := exec.Command("./a.sh")
	err := runCommand(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(cmd *exec.Cmd) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	go printOutputWithHeader(stdout, ansi.Color("stdout:", "green"))
	go printOutputWithHeader(stderr, ansi.Color("stderr:", "red"))

	return cmd.Wait()
}

func printOutputWithHeader(r io.Reader, header string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Printf("%s%s\n", header, scanner.Text())
	}
}
