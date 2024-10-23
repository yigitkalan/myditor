package terminal

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func EnableRaw(fd int) *unix.Termios {
	termios := GetTermios(fd)

	oldState := termios
	// Look manpage of termios(3) for more information about these flags
	termios.Lflag &^= unix.ECHO | unix.ICANON | unix.ISIG | unix.IEXTEN
	termios.Iflag &^= unix.IXON | unix.ICRNL | unix.BRKINT | unix.INPCK | unix.ISTRIP
	termios.Oflag &^= unix.OPOST
	termios.Cflag &^= unix.CS8
	termios.Cc[unix.VMIN] = 0
	termios.Cc[unix.VTIME] = 1

	err := unix.IoctlSetTermios(fd, unix.TCSETS, termios)
	if err != nil {
		Kill("Enabling Raw mode")
	}
	return oldState
}

func DisableRaw(fd int, original *unix.Termios) {
	err := unix.IoctlSetTermios(fd, unix.TCSETS, original)
	if err != nil {
		Kill("Disabling Raw mode")
	}
}

func GetTermios(fd int) *unix.Termios {
	term, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		Kill("Getting terminal information")
	}
	return term
}

func Kill(message string) {
	EditorRefreshScreen()
	fmt.Errorf(message)
	os.Exit(1)
}
