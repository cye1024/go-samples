package util

import (
	"fmt"
	"os"
	"runtime"
)

func HandleErr(e error) {
	if e != nil {
		_, f, l, _ := runtime.Caller(1)
		fmt.Printf("Err! %s:%d [%s] \n", f, l, e.Error())
	}
}

func NewFile(n string) (*os.File, error) {
	f, e := os.OpenFile(n, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if e == nil {
		return f, nil
	} else {
		return nil, e
	}
}
