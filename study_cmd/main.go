package main

import (
	"fmt"
	"log"
	"os/exec"
)

// ls -al | grep main
func main() {
	cmd1 := exec.Command("ls", "-al")
	pipout1, err := cmd1.StdoutPipe()
	if err != nil {
		log.Fatalf("%v", err)
	}

	cmd2 := exec.Command("grep", "main")
	cmd2.Stdin = pipout1

	cmd1.Start()
	defer cmd1.Wait()
	outb, err := cmd2.Output()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(string(outb))
}
