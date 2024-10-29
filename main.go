package main

import (
	"myditor/editor"
)

func main() {
    editor.Init()
	for {
		if editor.EditorProcessKey() == -1 {
			break
		}
	}
}
