package main

import (
	"bufio"
	"fmt"
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

	go func() {
		stdoutHeader := ansi.Color("stdout:", "green")
		stdoutScanner := bufio.NewScanner(stdout)
		for stdoutScanner.Scan() {
			fmt.Printf("%s%s\n", stdoutHeader, stdoutScanner.Text())
		}
	}()

	go func() {
		stderrHeader := ansi.Color("stderr:", "red")
		stderrScanner := bufio.NewScanner(stderr)
		for stderrScanner.Scan() {
			fmt.Printf("%s%s\n", stderrHeader, stderrScanner.Text())
		}
	}()

	return cmd.Wait()
}
