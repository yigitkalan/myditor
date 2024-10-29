package core

import "golang.org/x/sys/unix"

type EditorConfig struct {
    OriginalState unix.Termios 
    ScreenRows uint16
    ScreenCols uint16
}

var Config EditorConfig = EditorConfig{}
