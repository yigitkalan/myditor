package terminal

import (
	"fmt"
	"myditor/core"
	"os"
	"unicode"

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

func GetWindowSize(fd int, config *core.EditorConfig) {
	ws, _ := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	//Fallback way for unsupported systems
	// if err != nil {
	n, _ := os.Stdout.Write([]byte("\033[999C\033[999B"))
	if n != 12 {
		kill("Getting window size")
	}
	GetCursorPosition(int(config.ScreenRows), int(config.ScreenCols))
	// }
	if ws.Col == 0 {
		kill("Getting window size")
	}
	config.ScreenCols = ws.Col
	config.ScreenRows = ws.Row
}

func GetCursorPosition(rows int, columns int) {
	// Send the escape sequence to request cursor position
	os.Stdout.Write([]byte("\033[6n"))

	// Newline for clarity in output
	fmt.Print("\r\n")

	// Read and debug output of each byte
	b := make([]byte, 1)
	for {
		n, err := os.Stdin.Read(b)
		if err != nil || n == 0 {
			break
		}

		// Check if byte is a control character
		if unicode.IsControl(rune(b[0])) {
			fmt.Printf("%d\r\n", b[0])
		} else {
			fmt.Printf("%d ('%c')\r\n", b[0], b[0])
		}

		// Break when we reach 'R', which typically marks the end of the response
		if b[0] == 'R' {
			break
		}
	}
	EditorReadKey()
}

func kill(message string) {
	EditorRefreshScreen()
	DisableRaw(GetFd())
	fmt.Errorf(message)
	os.Exit(1)
}
