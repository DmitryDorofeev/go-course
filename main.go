package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	out, err := exec.Command("wc", "-l", os.Args[1]).Output()
	if err != nil {
		return
	}
	fmt.Printf("%s", out)
}
