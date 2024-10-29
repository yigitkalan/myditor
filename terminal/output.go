package terminal

import (
	"myditor/core"
	"os"
)


func EditorRefreshScreen() {
	clearScreen := "\033[2J"
	toHome := "\033[H"
	os.Stdout.Write([]byte(clearScreen + toHome))
    EditorDrawRows()
    os.Stdout.Write([]byte("\033[H"))
    os.Stdout.Write([]byte("\033[H"))
}

func EditorDrawRows() {
    for i := 0; i < int(core.Config.ScreenRows); i++ {
		os.Stdout.Write([]byte("~\r\n"))
	}
}
