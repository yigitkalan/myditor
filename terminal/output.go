package terminal

import "os"


func EditorRefreshScreen() {
	clearScreen := "\033[2J"
	toHome := "\033[H"
	os.Stdout.Write([]byte(clearScreen + toHome))
    EditorDrawRows()
    os.Stdout.Write([]byte("\033[H"))
}

func EditorDrawRows() {
	for i := 0; i < 24; i++ {
		os.Stdout.Write([]byte("~\r\n"))
	}
}
