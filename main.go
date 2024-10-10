package main

import (
	"myditor/terminal"
	"os"
)

func main(){
    fd := int(os.Stdin.Fd())
    original := terminal.EnableRaw(fd)
    terminal.LoopInput()
    defer terminal.DisableRaw(fd, original)
}
