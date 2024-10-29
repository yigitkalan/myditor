package editor

import (
	"myditor/core"
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

func Init(){
    terminal.SetWindowSize(terminal.GetFd(), &core.Config)
}

func exit(){
    terminal.EditorRefreshScreen()
    terminal.DisableRaw(terminal.GetFd())
}
