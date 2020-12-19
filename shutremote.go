package main

import (
	"fmt"
	"os/exec"
)

// https://github.com/lanastasov/windows_automation/blob/master/cmd-prompt/remote-shutdown-local-network.png
func main() {
	// shutremote 1 - execute first command
	// shutremote 2 - execute second command
	// shutremtoe 3 - executes first and second command
	
	// shutdown -s -m \\desktop-7huo5h1 -t 01
	// shutdown -s -m 192.168.1.14 -t 01
	cmd, err := exec.Command("shutdown", "/s", "/m", "\\\\desktop-7huo5h1", "/t", "01").Output()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("Command ran: %s\n", string(cmd))
}
