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

func main() {
	cmd := exec.Command("./a.sh")
	stdout, stderr, err := runCommand(cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("stdout result:%s\n", stdout)
	fmt.Printf("stderr result:%s\n", stderr)
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

	go printOutputWithHeader(outReader2, ansi.Color("stdout:", "green"))
	go printOutputWithHeader(errReader2, ansi.Color("stderr:", "red"))

	err = cmd.Wait()
	if err != nil {
		return
	}

	stdout = bufout.String()
	stderr = buferr.String()
	return
}

func printOutputWithHeader(r io.Reader, header string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Printf("%s%s\n", header, scanner.Text())
	}
}
