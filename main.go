package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("./a.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
