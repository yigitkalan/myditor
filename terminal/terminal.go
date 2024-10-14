package terminal

import (
	"fmt"
	"io"
	"os"
	"unicode"

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
		panic(err)
	}
	return oldState
}

func DisableRaw(fd int, original *unix.Termios) {
	err := unix.IoctlSetTermios(fd, unix.TCSETS, original)
	if err != nil {
		panic(err)
	}
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
		b := byte(0)
		b, err := readKey()
		if err != nil && err != io.EOF {
			break
		}
		if unicode.IsControl(rune(b)) {
			fmt.Printf("%d\n\r", b)
		} else {
			fmt.Printf("%d ('%c')\n\r", b, b)
		}
		if b == 'q' {
			break
		}
	}
}

// Read input byte by byte
func readKey() (byte, error) {
	buf := make([]byte, 1)
	n, err := os.Stdin.Read(buf)
	if n == 0 {
		return 0, err
	}
	return buf[0], err
}
