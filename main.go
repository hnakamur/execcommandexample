package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os/exec"
	"syscall"

	"github.com/mgutz/ansi"
)

const (
	stdoutColor = "green"
	stderrColor = "red"
)

func main() {
	var verbose = flag.Bool("v", false, "get verbose output")
	flag.Parse()

	cmd := exec.Command("./a.sh")
	stdout, stderr, exitCode, err := runCommand(cmd, *verbose)
	fmt.Printf("stdout result:%s\n", ansi.Color(stdout, stdoutColor))
	fmt.Printf("stderr result:%s\n", ansi.Color(stderr, stderrColor))
	fmt.Printf("exitCode:%d\n", exitCode)
	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(cmd *exec.Cmd, verbose bool) (stdout, stderr string, exitCode int, err error) {
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

	go printOutputWithHeader("stdout:", stdoutColor, outReader2, verbose)
	go printOutputWithHeader("stderr:", stderrColor, errReader2, verbose)

	err = cmd.Wait()

	stdout = bufout.String()
	stderr = buferr.String()

	if err != nil {
		if err2, ok := err.(*exec.ExitError); ok {
			if s, ok := err2.Sys().(syscall.WaitStatus); ok {
				err = nil
				exitCode = s.ExitStatus()
			}
		}
	}
	return
}

func printOutputWithHeader(header, color string, r io.Reader, verbose bool) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if verbose {
			fmt.Printf("%s%s\n", header, ansi.Color(scanner.Text(), color))
		}
	}
}
