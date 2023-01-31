package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("powershell", "-Command", "Get-Alias")
	output, err := cmd.CombinedOutput()
	if err != nil {

	}
	fmt.Println(output)

	var name = "test"
	var s = `hello`
	fmt.Println(s)
}
