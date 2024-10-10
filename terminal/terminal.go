package terminal

import (
	"fmt"
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
		if unicode.IsControl(rune(b)) {
            fmt.Printf("%d\n\r", b)
        } else{
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
	_, err := os.Stdin.Read(buf)
	return buf[0], err
}
