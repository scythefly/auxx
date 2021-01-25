package flock

import (
	"context"
	"fmt"
	"time"

	"github.com/kakami/flock"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flock",
		Short: "run flock examples",
	}

	cmd.AddCommand(
		newFlockCommand(),
		newTryFlockCommand(),
	)
	return cmd
}

func newFlockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lock",
		RunE: runFlock,
	}
	return cmd
}

func runFlock(_ *cobra.Command, _ []string) error {
	var err error
	fmt.Println("===================== run flock lock =====================")
	l := flock.NewDirLocker("./")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	locked, err := l.TryLockContext(ctx, time.Second)
	if err != nil {
		return errors.WithMessage(err, "try lock")
	}
	if !locked {
		return errors.New(">>> unlock")
	}
	fmt.Println(">>>> get locker")
	time.Sleep(20 * time.Second)
	fmt.Println(">>> release locker")
	l.Release()
	/*
		// fd, err := syscall.Open("lockfile", syscall.O_WRONLY|syscall.O_CREAT, 0644)
		file, err := os.OpenFile("lockfile", syscall.O_CREAT|syscall.O_RDWR|syscall.O_CLOEXEC, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		fst := syscall.Flock_t{
			Type:   syscall.F_WRLCK,
			Whence: io.SeekStart,
			Start:  0,
			Len:    0,
			Pid:    0,
		}
		ftt := syscall.Flock_t{
			Type: syscall.F_WRLCK,
		}
		fmt.Println("=====> syscall.FcntFlock/F_SETLKW")
		err = syscall.FcntlFlock(file.Fd(), syscall.F_SETLKW, &fst)
		if err != nil {
			return errors.WithMessage(err, "syscall.FcntlFlock/F_SETLKW")
		}
		fmt.Printf("%#v\n", fst)
		fmt.Println("=====> syscall.FcntFlock/F_GETLK")
		err = syscall.FcntlFlock(file.Fd(), syscall.F_GETLK, &ftt)
		if err != nil {
			return errors.WithMessage(err, "syscall.FcntlFlock/F_GETLK")
		}
		fmt.Printf("%#v\n", ftt)
		fmt.Println("owner?", ftt.Type == syscall.F_UNLCK)
		time.Sleep(10 * time.Second)
	*/
	return err
}

func newTryFlockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "try",
		RunE: runTryFlock,
	}
	return cmd
}

func runTryFlock(_ *cobra.Command, _ []string) error {
	var err error
	fmt.Println("===================== run flock tryLock =====================")
	/*
		fd, err := syscall.Open("lockfile", syscall.O_WRONLY|syscall.O_CREAT, 0644)
		if err != nil {
			return errors.WithMessage(err, "syscall.Open")
		}
		fgetlk := syscall.Flock_t{
			Type: syscall.F_WRLCK,
		}
		fmt.Println("=====> syscall.FcntFlock/F_GETLK")
		err = syscall.FcntlFlock(uintptr(fd), syscall.F_GETLK, &fgetlk)
		if err != nil {
			return errors.WithMessage(err, "syscall.FcntlFlock/F_GETLK")
		}
		fmt.Printf("%#v\n", fgetlk)
	*/
	return err
}
