package samples

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	. "fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/davecgh/go-spew/spew"
	jsoniter "github.com/json-iterator/go"
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
	//var (
	//a float64 = 1.234563224
	//b float64 = 1.234563224
	//)
	//Printf("%t, %8.8f %8.8f\n", a+0.004 == b+0.004, a+0.004, b+0.004)
	//for i := 0; i < 100000; i++ {
	//a += 0.01
	//b += 0.01
	//}
	//Printf("%t, %8.8f %8.8f\n", a == b, a, b)

	var c float64 = 0
	var d float64 = 100

	for i := 0; i < 5000; i++ {
		c += .01
	}
	for i := 0; i < 5000; i++ {
		d -= .01
	}
	Println(c, d, c == d, IsEqual(c, d), IsEqual(d, c))
}

func T19() {
	s := "CCPUID_704=xxxxxxxxxxxx; CCPUID_707=aaaaa; Idea-69579d97=9e592d64-3dce-4400-94bb-020e23b90c4b; V=5911167682205782503"
	Println(strings.Index(s, "CCPUID_707="))
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

const MIN = 0.000001

func IsEqual(a, b float64) bool {
	return math.Dim(a, b) < MIN
}
func T20() {
	t := time.Now()

	f, e := os.OpenFile("/Users/cye/tmp/t.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if e != nil {
		log.Panic(e)
	}

	var m sync.Mutex
	s := []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaa\n", "bbbbbbbbbbbbbbbbbbbbbbbbbb\n", "ccccccccccccccccccccccc\n"}
	for i := 0; i < 100000; i++ {
		go func() {
			for _, v := range s {
				go func(s string) {
					m.Lock()
					defer m.Unlock()
					n, e := f.WriteString(s)
					if e != nil {
						log.Panic(n, e)
					}
				}(v)
			}
		}()
	}
	Println(time.Since(t))
}

func T21() {
	t := time.Now()
	var w *bufio.Writer

	chs := make(chan string)
	f, e := os.OpenFile("/Users/cye/tmp/t.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if e != nil {
		log.Panic(e)
	}
	//w = bufio.NewWriter(f)
	w = bufio.NewWriterSize(f, 9500)

	//var m sync.Mutex
	go func() {
		w.WriteString(<-chs)
	}()
	s := []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaa\n", "bbbbbbbbbbbbbbbbbbbbbbbbbb\n", "ccccccccccccccccccccccc\n"}
	for i := 0; i < 100000; i++ {
		go func() {
			for _, v := range s {
				go func(s string) {
					//m.Lock()
					//defer m.Unlock()
					chs <- s
					//n, e := w.WriteString(s)
					//n, e := f.Write([]byte(s))
					if e != nil {
						log.Panic(e)
					}
				}(v)
			}
		}()
	}
	w.Flush()
	Println(time.Since(t))
}

func T22() {
	m := make(map[string]interface{})
	m["one"] = 1
	if v, ok := m["one"].(int); ok {
		Println(v, ok)
	}
	var v interface{}
	_ = v.(string)
	//if a, ok := v.(string); ok {
	//}
}

func T23() {
	//b1 := bytes.NewBuffer(nil)
	//b2 := bufio.NewReadWriter(b1)
	//net.Conn
}

func T24() {
	type t struct {
		s string
	}
	m := make(map[int]*t)
	m[1] = &t{s: "sss"}
	Printf("%+v \n", m[1].s)
	for _, v := range m {
		func(x *t) {
			y := x
			y.s = "xxx1"
		}(v)
	}

	Printf("%+v \n", m[1].s)
}

func T25() {
	type x struct {
		s string
		i int
	}

	type t struct {
		f func(*x)
	}

	tf := func(i int, s string) *t {
		return &t{
			func(x *x) {
				x.s = s
				x.i = i
			},
		}
	}

	tt := tf(3, "three")

	tx := new(x)
	tt.f(tx)
	Println(tx)
}

func T26() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		wg.Done()
		go func() {
			time.Sleep(time.Second)
			Println("hahah")
			wg.Done()
		}()
		Println("end1")
	}()

	wg.Wait()
	Println("end")
}

func T27() {
	//rune.New()
	//byte.Rune()
	//bytes.Runes()
}

func T28() {
	//bm, err := cache.NewCache()
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	logs.Debug("Hello beeog log!")
	logs.Info("info")
	logs.Warn("warn")
	logs.Warning("warning")
	logs.Error("error")

	log.SetFlags(8)
	log.Println("hello log")
	log.Panic("aaa")
}

func T29() {
	//name, e := os.Hostname()
	//if e != nil {
	//logs.Error(e)
	//}

	//addrs, e := net.LookupHost(name)
	//if e != nil {
	//logs.Error(e)
	//}

	//addrs, e := net.InterfaceAddrs()
	//if e != nil {
	//logs.Error(e)
	//}

	//ifaces, _ := net.Interfaces()
	//// handle err
	//for _, i := range ifaces {
	//addrs, _ := i.Addrs()
	//// handle err
	//for _, addr := range addrs {
	//var ip net.IP
	//switch v := addr.(type) {
	//case *net.IPNet:
	//ip = v.IP
	//if ip.To4() != nil && ip.String() != "127.0.0.1" {
	//logs.Debug(ip.String())
	//}
	//case *net.IPAddr:
	//ip = v.IP
	//}
	//}

	//}

	logs.Debug(GetLocalIP())
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""

	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()

			}

		}

	}
	return ""

}

func T30() {
	m := make(map[string]*int64)
	var i int64 = 1
	m["one"] = &i

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(m["one"], int64(1))
		}()
	}
	wg.Wait()
	Println(*m["one"])
}

func T31() {
	m := make(map[string]*int64)
	var i int64 = 1
	m["one"] = &i

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			atomic.AddInt64(m["one"], int64(1))
			x := strconv.Itoa(i % 10)
			if _, ok := m[x]; ok {
				atomic.AddInt64(m[x], int64(1))
			} else {
				y := int64(i)
				m[x] = &y
			}
		}(i)
	}
	wg.Wait()
	for k, v := range m {
		Println(k, *v)
	}
}

func T32() {
	m := make(map[string]*int64)

	var l sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			x := strconv.Itoa(i % 10)
			l.RLock()
			if _, ok := m[x]; ok {
				atomic.AddInt64(m[x], int64(1))
			} else {
				l.Lock()
				y := int64(0)
				m[x] = &y
				l.Unlock()
			}
			l.RUnlock()
		}(i)
	}
	wg.Wait()
	for k, v := range m {
		Println("a", k, *v)
	}
}

type t struct {
	p int
}

type it interface {
	setp(int)
}

func (t t) setp(p int) {
	t.p = p
}

func (t t) getp() int {
	return t.p
}

func T33() {
	t := new(t)
	t.setp(4)

	t1 := *t
	t1.setp(9)

	t2 := t
	t2.setp(19)
	Println(t.p, t1.p)
}

func T34() {
	type t struct {
		i int
	}

	m := make(map[*t]string)
	s := &t{3}
	Println(s)

	m[s] = "aaa"
	Println(m)

	s.i = 5
	Println(m)
}

func T35() {
	type t struct {
		i int
	}

	s := t{3}
	m := t{3}
	Println(s == m) // true

	s1 := &t{3}
	m1 := &t{3}
	Println(s1 == m1) // false

	s2 := &t{3}
	m2 := s2
	Println(s2 == m2) // true
}

func T36() {
	type t struct {
		i int
	}

	m := make(map[t]string)

	s := t{1}
	m[s] = "one"

	m[s] = "two"
}

func T37() {
	type t struct {
		i int
	}

	m := make(map[int]t)
	m[1] = t{2}
	//m[1].i = 3
	Println(m)
}

func T38() {
	switch "1" {
	case "1":
		Println("1")
		fallthrough
	case "2":
		Println("2")
		Println("2")
	}
}

func T39() {
	t := new(t)
	var i it
	i = t
	i.setp(1)
}

var jsonstr string = `
{
	"id": "a01-04rH-02FSRH-0_B-1IP",
    "app": {
        "content": {
            "keywords": "",
            "title": "xxx160220",
            "ext": {
                "channel": 31,
                "usr": "113077370",
                "cs": "110319999",
                "s": "252245"
            }
        },
        "name": "xxx"
    }
}
`

type jsonStruct struct {
	Id  string
	App *app
}

type app struct {
	Content *content
	Name    string
}

type content struct {
	Keywords string
	Title    string
	Ext      *ext
}

type ext struct {
	Channel int    `json:"channel"`
	Usr     string `json:"usr"`
	Cs      string `json:"cs"`
	Vid     string `jsong:"vid"`
	S       string `json:"s"`
}

func T40() {
	j := new(jsonStruct)
	e := json.Unmarshal([]byte(jsonstr), j)
	if e != nil {
		Println(e)
	}
	//spew.Dump(j)
}

func T41() {
	j := new(jsonStruct)
	e := jsoniter.Unmarshal([]byte(jsonstr), j)
	if e != nil {
		Println(e)
	}
	spew.Dump(j)
}

func T42() {
	r1 := []int{1, 2, 3, 5, 7}
	r2 := []int{1, 3, 2, 5, 8}
loopppp:
	for i, j := range r1 {
		for k, v := range r2 {
			if j == v {
				Println(j, v)
				r1 = append(r1[:i], r1[(i+1):]...)
				r2 = append(r2[:k], r2[(k+1):]...)
				goto loopppp
			}
		}
	}

	if len(r1) == 0 && len(r2) == 0 {
		Println("OK")
	}
}

func T43() {
	r1 := []int{1, 2, 3, 5, 7}
	r2 := []int{1, 3, 2, 5, 8}
	si1 := sort.IntSlice(r1)
	si2 := sort.IntSlice(r2)
	si1.Sort()
	si2.Sort()
	Println(si1, si2)
}

func T44() {
	type T struct {
		a *string
	}

	s := "aaa"
	t := &T{}
	t.a = &s

	t1 := *t
	Printf("%p %p \n", t, &t1)
}

func T45() {
	type t1 struct {
		i int64
	}

	t := &t1{
		i: 3000,
	}
	var w sync.WaitGroup
	w.Add(2500)
	for k := 0; k < 2500; k++ {
		go func() {
			atomic.AddInt64(&t.i, -1)
			w.Done()
		}()
	}

	w.Wait()
	Println(t.i)
}
