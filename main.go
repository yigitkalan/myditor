package main

import (
	"myditor/editor"
	"myditor/terminal"
)

func main() {
	fd := int(terminal.GetFd())
	terminal.EnableRaw(fd)
	terminal.EditorRefreshScreen()
	for {
		if editor.EditorProcessKey() == -1 {
			break
		}
	}
}
