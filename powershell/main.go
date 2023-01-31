package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

type PowerShell struct {
	ctx      context.Context
	cmd      *exec.Cmd
	pwshPath string
	pid      string
}

func NewPowerShell() *PowerShell {
	ps, _ := exec.LookPath("pwsh")
	return &PowerShell{
		pwshPath: ps,
	}
}

//func (p *PowerShell) execute(args ...string) (stdOut string, stdErr string, err error) {
//
//	args = append([]string{"-ExecutionPolicy", "Bypass", "-NoExit", "-Command"}, args...)
//	cmd := exec.Command("pwsh", args...)
//
//	var stdout bytes.Buffer
//	var stderr bytes.Buffer
//	cmd.Stdout = &stdout
//	cmd.Stderr = &stderr
//	err = cmd.Start()
//	fmt.Println(cmd.Process.Pid)
//	stdOut, stdErr = stdout.String(), stderr.String()
//	return
//}

func (c *PowerShell) execute(args ...string) (stdOut string, stdErr string, err error) {
	//args = append([]string{"-ExecutionPolicy", "Bypass", "-NoExit", "-Command"}, args...)

	c.cmd = exec.CommandContext(c.ctx, c.pwshPath, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.cmd.Stdout = &stdout
	c.cmd.Stderr = &stderr
	err = c.cmd.Run()
	fmt.Println(c.cmd.Process.Pid)
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

//func MakePowerShellSession() (session PowerShell, err error) {
//	ps, _ := exec.LookPath("pwsh")
//	session := PowerShell{cmd: exec.CommandContext()}
//	err = c.cmd.Start()
//	fmt.Println(c.cmd.Process.Pid)
//	return
//}

func main() {
	//ctx, _ := context.WithCancel(context.Background())
	//pwshPath, _ := exec.LookPath("pwsh")
	//
	//pSession := PowerShell{
	//	ctx:      ctx,
	//	cmd:      exec.CommandContext(ctx, "pwsh"),
	//	pwshPath: pwshPath,
	//}
	//pSession.execute("Get-Alias")

	test := exec.Command("pwsh", "Get-Alias")
	output, err := test.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(output))

	//test = exec.Command("Get-Alias")

	//output, err = test.Output()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("PS Get-Alias")
	//fmt.Println(string(output))

	lsCmd := exec.Command("echo", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

//func main() {
//	pw := NewPowerShell()
//	execute, _, err := pw.execute("Get-Alias")
//	if err != nil {
//		return
//	}
//
//	fmt.Println(execute)
//}

//func main() {
//	// choose a backend
//	back := &backend.Local{}
//
//	// start a local powershell process
//	shell, err := ps.New(back)
//	if err != nil {
//		panic(err)
//	}
//	defer shell.Exit()
//
//	// ... and interact with it
//	stdout, _, err := shell.Execute("Get-WmiObject -Class Win32_Processor")
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(stdout)
//}

//
//func main() {
//
//	// run a command
//	gosh.PowershellCommand(`Write-Host "hello from powershell"`)
//
//	// run a script
//	gosh.PowershellCommand(`
//  $git_username = git config user.name
//
//  Write-Host $git_username
//`)
//
//	// run a command with output
//	err, out, errout := gosh.PowershellOutput(`[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$APP_PATH\bin", [System.EnvironmentVariableTarget]::User)`)
//
//	if err != nil {
//		log.Printf("error: %v\n", err)
//		fmt.Print(errout)
//	}
//
//	fmt.Print(out)
//
//}
