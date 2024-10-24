package core

import "golang.org/x/sys/unix"

type EditorConfig struct {
    OriginalState unix.Termios 
}

var Config EditorConfig = EditorConfig{}
