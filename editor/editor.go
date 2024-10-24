package editor

import (
	"myditor/terminal"
	"os"
)


func EditorProcessKey() int  {
	b := terminal.EditorReadKey()

	switch b {
	case terminal.CTRL_KEY('q'):
        exit()
        return -1
    default:
        os.Stdout.Write([]byte{b})
	}
    return 0
}

func exit(){
    terminal.EditorRefreshScreen()
    terminal.DisableRaw(terminal.GetFd())
}
