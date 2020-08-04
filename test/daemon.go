package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func Daemon() {
	var cmd *exec.Cmd
	if len(os.Args) > 1 {
		cmd = exec.Command(os.Args[0], os.Args[1:]...)
	} else {
		cmd = exec.Command(os.Args[0])
	}
	f, err := os.Create(filepath.Join(os.Args[0], "../out.log"))
	if err != nil {
		panic(err)
	}
	cmd.Stdin = nil
	cmd.Stdout = f
	cmd.Stderr = f
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	err = cmd.Start()
	if err == nil {
		cmd.Process.Release()
		os.Exit(0)
	}
}
