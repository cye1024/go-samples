package samples

import (
	"bytes"
	"fmt"
)

func NT() {
	//c, e := net.Dial("ip", ":8084")
	b := bytes.NewBuffer(nil)
	c := bytes.NewBuffer(nil)
	b.WriteString("aaa")
	b.WriteString("bbb")
	fmt.Println(b.Len(), b.Cap(), b.String(), c.String())
}
