package samples

import (
	. "fmt"
	. "util"
)

func T_io() {
	f, e := NewFile("/Users/cye/tmp/hello.log")
	HandleErr(e)
	_, e = f.Write([]byte("aaaabb"))
	HandleErr(e)
	info, e := f.Stat()
	HandleErr(e)
	Println(info.Size())
	var b = make([]byte, 100)
	_, e = f.Read(b)
	if e.Error() == "EOF" {
		Println("EOR")
	} else {
		HandleErr(e)
	}

	Println(string(b))
}
