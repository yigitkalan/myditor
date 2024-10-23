package terminal

import "os"


func EditorRefreshScreen() {
	clearScreen := "\033[2J"
	toHome := "\033[H"
	os.Stdout.Write([]byte(clearScreen + toHome))
}

func EditorDrawRows(){
    
}

