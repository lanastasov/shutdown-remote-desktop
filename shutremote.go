package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// shutdown -r -m \\desktop-7huo5h1 -t 01
	cmd, err := exec.Command("shutdown", "/s", "/m", "\\\\desktop-7huo5h1", "/t", "01").Output()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("Command ran: %s\n", string(cmd))
}
