package samples

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	. "fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func T1() {
	var a int

	a = 2
	Printf("%p \n", &a)
	a, b := 4, "bb"
	Printf("%p %v, %v \n", &a, a, b)
}

func T2() {
	var s [2][3]int
	for i, v := range s {
		Println(i, v)
	}
	Println(len(s))
}

func T3() {
	errors.New("aa")
	s := make([]interface{}, 3, 5)
	s[2] = 3
	s[1] = 3

	Println(s)
	func() {
		return
	}()
}

type testint interface {
	Strs1() []string
}

type I struct {
	val int
	str []string
}

func (i I) Value() int {
	return i.val
}

func (i I) Strs() []string {
	return i.str
}
func (i *I) Strs1() []string {
	return i.str
}
func T4() {
	//var i1 testint = I{}
	//i1.Strs1()
	var i2 testint = &I{}
	i2.Strs1()
}

func T5() {
	Println("==", runtime.NumCPU(), runtime.GOMAXPROCS(2218))
	ch := make(chan int, 5)
	for i := 0; i < 10; i++ {
		go func(x int) {
			ch <- x
		}(i)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			ch <- 5
		}
	}()

	for i := range ch {
		Println(i)
	}
}

func T6() {
	s := "aaa!=bbb!ccc>=ddd"
	s = strings.Replace(s, "!=", "~", -1)
	Println(s)
	s = strings.Replace(s, "~", "!=", -1)
	Println(s)
}

func T7() {
	Nobid(23, "aa")
	Nobid()
}

func (i *I) String() string {
	return strings.Join(i.str, ", ")
}

func TestString() {
	i := new(I)
	i.str = []string{"aa", "bb"}
	Println(i)
	return
}

func T8() {
	TestString()
}

func T9() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	Println("type:", p.Type())
	Printf("Canset? %t, %v \n", p.CanSet(), p)

	v := p.Elem()

	Println("Canset?:", v.CanSet())
	v.SetFloat(4.5)
	Println(v.Interface())
	Println(x)
}

func T10() {
	Println("debug", "string", "ok")
	Printf("debug "+"string %s\n", "ok")

	t := time.Now()
	defer func(t *time.Time) {
		Println(time.Since(*t))
	}(&t)
	time.Sleep(time.Second)
}

func T11() {
	//bufio.Reader()
}

func T12() {
	cmd0 := exec.Command("ls", "-l", "/Users/cye")

	outPipe0, e := cmd0.StdoutPipe()
	if e != nil {
		Println(e)
	}
	e = cmd0.Start()
	if e != nil {
		Println(e)
	}
	ret := bytes.NewBuffer(nil)
	for {
		b := make([]byte, 100)
		n, e := outPipe0.Read(b)
		if e != nil {
			if e == io.EOF {
				break
			}
			Println(e)
			return
		}
		if n > 0 {
			ret.Write(b[:n])
		}
	}
	Println(ret.String())
}

func T13() {
	cmd := exec.Command("ls", "-l", "/Users/cye")
	p, _ := cmd.StdoutPipe()
	b := bufio.NewReader(p)
	cmd.Start()
	ret := bytes.NewBuffer(nil)
	for {
		//l, e := b.ReadSlice('\n')
		l, _, e := b.ReadLine()
		if e != nil {
			if e == io.EOF {
				break
			}
			Println(e)
			return
		}
		ret.Write(l)
		ret.WriteString("\n")
	}
	Println(ret.String())
}

func T14() {
	for i := 0; i < 10; i++ {
		Println(GetRandomString(8))
	}
}

func T15() {
	sigRec := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGINT}
	signal.Notify(sigRec, sigs...)
	var i int
	for sig := range sigRec {
		Println("recieve:", sig)
		i++
		if i > 3 {
			break
		}
	}

	signal.Stop(sigRec)
	close(sigRec)

	x := os.Stdin
	b := bufio.NewReader(x)
	Println(b.ReadString('\n'))
}

func T16() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	var a, b, c int
	for i := 0; i < 10000; i++ {
		for k, _ := range m {
			switch k {
			case "one":
				a++
			case "two":
				b++
			case "three":
				c++
			}
			break
		}
	}

	Println(a, b, c)
}

func T17() {
	for i := 0; i < 5; i++ {
		Println(Random(15, 20))
	}
}

func T18() {
	var (
		a float64 = 1.23456
		b float64 = 1.23456
	)
	Printf("%t, %8.8f\n", a == b, a)
}

func Random(args ...int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	switch len(args) {
	case 0:
		return r.Int()
	case 1:
		return r.Intn(args[0])
	case 2:
		return r.Intn(args[1]-args[0]) + args[0]
	default:
		return -1
	}
}

func md5IfNonEmpty(s string) string {
	if s == "" {
		return ""
	}
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])

}

//TODO: move to utils
func GetRandomString(count int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])

	}
	return string(result)

}

func Nobid(args ...interface{}) {
	Println(args...)
}
