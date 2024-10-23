package main

import (
	"myditor/terminal"
	"os"
)

func main() {
	fd := int(os.Stdin.Fd())
	original := terminal.EnableRaw(fd)
	defer terminal.DisableRaw(fd, original)
	terminal.EditorRefreshScreen()
	for {
		terminal.EditorProcessKey()
	}
}
