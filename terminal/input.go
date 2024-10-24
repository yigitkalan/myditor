package terminal

import (
	"io"
	"os"
)


// Read input byte by byte
func EditorReadKey() byte {
	buf := make([]byte, 1)
	var n int = 0
	var err error
    //this will stop the program until a key is pressed
	for n != 1 {
		n, err = os.Stdin.Read(buf)
	}

	if err != nil && err != io.EOF {
        kill("Reading keys")
	}
	return buf[0]
}

func CTRL_KEY(k byte) byte {
	return k & 0x1f
}
