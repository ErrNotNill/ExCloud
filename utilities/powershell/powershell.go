package powershell

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func PwShell() {
	posh := New()
	enableTimer := fmt.Sprintf("%s\n%s", sleeptimer10sec)
	stdOut, stdErr, err := posh.execute(enableTimer)
	fmt.Printf("\nEnableTimer:\nStdOut : '%s'\nStdErr: '%s'\nErr: %s", strings.TrimSpace(stdOut), stdErr, err)
	time.Sleep(time.Second * 30)
	stopTimer := fmt.Sprintf("%s\n%s", stopTimer)
	stdOut, stdErr, err = posh.execute(stopTimer)
	fmt.Printf("\nEnableTimer:\nStdOut : '%s'\nStdErr: '%s'\nErr: %s", strings.TrimSpace(stdOut), stdErr, err)
}

type PowerShell struct {
	powerShell string
}

// New create new session
func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

var (
	sleeptimer10sec = `shutdown -s -t 300`
	stopTimer       = `shutdown /a`
)
