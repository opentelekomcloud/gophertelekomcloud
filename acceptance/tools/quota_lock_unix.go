//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package tools

import (
	"os"

	"golang.org/x/sys/unix"
)

func tryLockFile(f *os.File) error {
	return unix.Flock(int(f.Fd()), unix.LOCK_EX|unix.LOCK_NB)
}

func unlockFile(f *os.File) error {
	return unix.Flock(int(f.Fd()), unix.LOCK_UN)
}
