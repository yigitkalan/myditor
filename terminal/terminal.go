package terminal

import (
	"fmt"
	"myditor/core"
	"os"
	"golang.org/x/sys/unix"
)


func EnableRaw(fd int) {
	termios := GetTermios(fd)

    core.Config.OriginalState = *termios

	// Look manpage of termios(3) for more information about these flags
	termios.Lflag &^= unix.ECHO | unix.ICANON | unix.ISIG | unix.IEXTEN
	termios.Iflag &^= unix.IXON | unix.ICRNL | unix.BRKINT | unix.INPCK | unix.ISTRIP
	termios.Oflag &^= unix.OPOST
	termios.Cflag &^= unix.CS8
	termios.Cc[unix.VMIN] = 0
	termios.Cc[unix.VTIME] = 1

	err := unix.IoctlSetTermios(fd, unix.TCSETS, termios)
	if err != nil {
		kill("Enabling Raw mode")
	}
}

func DisableRaw(fd int) {
	err := unix.IoctlSetTermios(fd, unix.TCSETS, &core.Config.OriginalState)
	if err != nil {
		kill("Disabling Raw mode")
	}
}

func GetTermios(fd int) *unix.Termios {
	term, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		kill("Getting terminal information")
	}
	return term
}

func GetFd() int {
    return int(os.Stdin.Fd())
}

func kill(message string) {
	EditorRefreshScreen()
    DisableRaw(GetFd())
	fmt.Errorf(message)
	os.Exit(1)
}

