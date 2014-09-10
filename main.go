package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/mgutz/ansi"
)

const (
	stdoutColor = "green"
	stderrColor = "red"
)

func main() {
	cmd := exec.Command("./a.sh")
	stdout, stderr, err := runCommand(cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("stdout result:%s\n", ansi.Color(stdout, stdoutColor))
	fmt.Printf("stderr result:%s\n", ansi.Color(stderr, stderrColor))
}

func runCommand(cmd *exec.Cmd) (stdout, stderr string, err error) {
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	var bufout, buferr bytes.Buffer
	outReader2 := io.TeeReader(outReader, &bufout)
	errReader2 := io.TeeReader(errReader, &buferr)

	if err = cmd.Start(); err != nil {
		return
	}

	go printOutputWithHeader("stdout:", stdoutColor, outReader2)
	go printOutputWithHeader("stderr:", stderrColor, errReader2)

	err = cmd.Wait()
	if err != nil {
		return
	}

	stdout = bufout.String()
	stderr = buferr.String()
	return
}

func printOutputWithHeader(header, color string, r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Printf("%s%s\n", header, ansi.Color(scanner.Text(), color))
	}
}
