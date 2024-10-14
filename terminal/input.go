package terminal

import (
	"io"
	"os"
)

func EditorProcessKey() {
	b := editorReadKey()

	switch b {
	case CTRL_KEY('q'):
        EditorRefreshScreen()
		os.Exit(0)
        break
    default:
        os.Stdout.Write([]byte{b})
	}
}

// Read input byte by byte
func editorReadKey() byte {
	buf := make([]byte, 1)
	var n int = 0
	var err error
    //this will stop the program until a key is pressed
	for n != 1 {
		n, err = os.Stdin.Read(buf)
	}

	if err != nil && err != io.EOF {
        Kill("Reading keys")
	}
	return buf[0]
}

func CTRL_KEY(k byte) byte {
	return k & 0x1f
}
