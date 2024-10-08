package terminal

import (
	"os"
	"golang.org/x/sys/unix"
)


func EnableRaw(fd int) *unix.Termios {
	termios := GetTermios(fd)

	oldState := termios
	termios.Lflag &^= unix.ECHO

	unix.IoctlSetTermios(fd, unix.TCSETS, termios)
	return oldState
}

func DisableRaw(fd int, original *unix.Termios) {
	unix.IoctlSetTermios(fd, unix.TCSETS, original)
}


func GetTermios(fd int) *unix.Termios {
	term, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		panic(err)
	}
	return term
}

func LoopInput() {
	for {
		b, err := readKey()
		if err != nil {
			break
		}
		if b == 'q' {
			break
		}
	}
}

// Read input byte by byte
func readKey() (byte, error) {
	buf := make([]byte, 1)
	_, err := os.Stdin.Read(buf)
	return buf[0], err
}
