package samples

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/binary"
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
	"unsafe"

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

func T46() {
	type t struct {
		i int
	}

	k := new(t)
	k.i = 1

	j := unsafe.Pointer(&k)
	g := (*unsafe.Pointer)(j)
	spew.Dump(j, k, g)
}

func T47() {
	type t struct {
		Data map[string]int
	}

	x := new(t)

	s := `{"data":{"aaa":3,"bbb":3}}`

	e := jsoniter.Unmarshal([]byte(s), x)
	if e != nil {
		Println(e)
		return
	}
	spew.Dump(x)
	return
}

func T48() {
	//  open file
	//file, err := os.OpenFile("file.log", os.O_RDONLY, 0600)
	//if err != nil {
	//log.Fatalln(err)

	//}

	//lines := make([]string, 10000)
	//files := make([]bufio.Reader, 0)
	//buff := bufio.NewReader(file)
	//for i := 0; ; i++ {
	//line, is, err := buff.ReadLine()
	//if err != nil && err != io.EOF {
	//// handle err
	//}
	//lines[i] = line

	//// file end
	//if err == io.EOF || !is {
	//break
	//}

	//// write to file
	//if i%10000 == 0 {
	//sort.Strings(lines)
	//// write file
	//f, e := os.Create(time.Now().String())
	//if e != nil {
	//// handle err
	//}
	//for _, l := range lines {
	//f.WriteString(l)
	//}
	//lines = lines[:0]
	//files = append(files, bufio.NewReader(f))
	//}
	//}

	//outFile := os.Create("outfile")
	//// read every file first number to compare
	//for {

	//ls := make([]string, 0, len(files))
	//for _, f := range files {
	//line, is, err := f.ReadLine()
	//if err == io.EOF || !is {
	//continue
	//}
	//ls = append(ls, line)
	//}
	//sort.Strings(ls)

	//if len(ls) > 0 {
	//outFile.Write([]byte(ls[0]))
	//outFile.Write('\n')
	//}
	//}
}

func T49() {
	str := []string{"1", "32", "12", "5"}
	itr := make([]int, len(str))
	for i, v := range str {
		ai, e := strconv.Atoi(v)
		if e != nil {
			//handle err
			continue
		}

		itr[i] = ai
	}

	sort.Ints(itr)
	Println(itr)
}

type adserver struct {
	ip    string
	timer *time.Timer
}

func (a *adserver) Start() {
	select {
	case <-a.timer.C:
		Println(a.ip)
	}
}

func T50() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			Println(i)
		}
	}()

	m := make(map[string]*adserver)
	a := new(adserver)
	a.ip = "111"
	a.timer = time.NewTimer(time.Second * 3)
	go a.Start()

	a1 := new(adserver)
	a1.ip = "112"
	a1.timer = time.NewTimer(time.Second * 5)
	go a1.Start()

	m[a.ip] = a
	m[a1.ip] = a1

	time.Sleep(time.Second * 2)
	if v, ok := m["111"]; ok {
		Println("reset")
		v.timer.Reset(time.Second * 3)
	}

	time.Sleep(time.Second * 10)
	//select {
	//case <-a.timer.C:
	//Println(a.ip)
	//case <-a1.timer.C:
	//Println(a1.ip)
	//}
}

func T51() {
	defer Println("defer")
	for {
		Println("hahah")
		time.Sleep(time.Second)
	}
}

func T52() {
	type t struct {
		i  int
		s  string
		ss []string
	}

	var t1 = new(t)
	t2 := *t1
	spew.Dump(t1, t2)
}

func T53() {
	c := make(chan string)
	go ping(c)
	go print(c)

	var input string
	Scanln(&input)
}

func ping(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- strconv.Itoa(i)
	}
}

func print(c <-chan string) {
	for {
		Println("receving:", <-c)
		time.Sleep(time.Second)
	}
}

func T54() {
	s := "abcdefg"
	Println(s[1:3])
}

type t54 struct {
	id string
	n  int
}

type tSli []*t54

func (t tSli) Len() int {
	return len(t)
}

func (t tSli) Less(i, j int) bool {
	return t[i].n < t[j].n
}

func (t tSli) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func T55() {
	//t := tSli{&t54{"1", 3}, &t54{"2", 5}, &t54{"3", 2}}
	t := make(tSli, 0)
	t = append(t, &t54{"4", 1})
	t = append(t, &t54{"1", 4})
	t = append(t, &t54{"2", 3})
	spew.Dump(t)
	sort.Sort(t)
	spew.Dump(t)
}

func T56() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case n := <-ch:
			Println(n)
		case ch <- i:
		}
	}
}

func T57() {
	type T struct {
		i int
		_ struct{}
		j int32
	}

	t := new(T)
	Println(unsafe.Sizeof(*t))

	Printf("%p %p %p", t, &t.i, &t.j)
}

func T58() {
	type T struct {
		A byte  `json:"a"`
		B int16 `json:"b"`
		C int32 `json:"c"`
		D []int `json:"d"`
	}

	t := T{}
	typ := reflect.TypeOf(t)

	n := typ.NumField()
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		Println(f.Name, f.Offset, f.Type.Size(), f.Type.Align(), f.Type.Kind())
	}
}

func T59() {
	ch := make(chan int, 1)
	ch <- 5
	Println(<-ch)
	close(ch)

	ch1 := make(chan int, 0)
	ch1 <- 5
	Println(<-ch1)
}

func T60() {
	m := make([]int, 0, 10)
	Printf("%p \n", m)
	for i := 0; i < 10; i++ {
		m = append(m, 3)
		Printf("%d : %p \n", i, &m[i])
	}
	i := 3
	Printf("i : %p \n", &i)
}

func T61() {
	logs.Debug("debug")
	logs.Warn("warn")
	logs.Warning("Warning")
	logs.Info("info")
	logs.Error("error")
}

func send(c chan<- int) {
	i := 0
	for {
		i++
		if i == 10 {
			break
		}
		Println("Send before", i)
		c <- i
		Println("                  Send after", i)
	}
}

func recieve(c <-chan int) {
	var i int
	for {
		Println("                  Recieve before", i)
		i = <-c
		Println("Recieve after", i)
	}
}

func T62() {
	ch := make(chan int)
	go send(ch)
	go recieve(ch)

	var s string
	Scanln(&s)
}

func T63() {
	const (
		a = 3
		b
		c = iota
		d
		e = 5
		f
		g = iota
	)

	Println(a, b, c, d, e, f, g)
}

func T64() {
	m := make(map[string]interface{})
	s := `{"age":5, "name":"goudan"}`

	e := jsoniter.UnmarshalFromString(s, &m)
	//e := json.Unmarshal([]byte(s), &m)
	if e != nil {
		Println(e)
		return
	}
	spew.Dump(m)
}

func T65() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)

	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	data := []byte(`{"Hello":{"Name":"leaf","Content":"Nice to meet you!"}}`)

	// len + data
	m := make([]byte, 2)
	m1 := make([]byte, 2)

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))
	binary.LittleEndian.PutUint16(m1, uint16(len(data)))
	Println(m, "|", m1)

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)
	for i, v := range m {
		Println(i, string(v))
	}

	time.Sleep(time.Second * 5)
}

func T66() {
	//var stringBuf = sync.Pool{
	//New: func() interface{} {
	//return nil
	//},
	//}
}

func T67() {
	var s string
	for i := 0; i < 1000; i++ {
		s = s + "a"
	}
}

func T68() {
	var s = bytes.NewBuffer(nil)
	for i := 0; i < 1000; i++ {
		s.WriteString("a")
	}
}

func T69() {
	m := &sync.Map{}
	var w sync.WaitGroup
	var n int

	w.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer w.Done()
			m.Store("one", 1)

			if v, ok := m.Load("one"); ok {
				n += v.(int)
			}
		}()
	}
	w.Wait()
	Println(n)
}

func T70() {
	m := make(map[string]int)
	var w sync.WaitGroup
	var mu sync.Mutex
	var n int

	w.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer w.Done()

			mu.Lock()
			defer mu.Unlock()
			m["one"] = 1

			if j, ok := m["one"]; ok {
				n += j
			}
		}()

	}
	w.Wait()
	Println(n)
}

func T71() {
	x := 1
	y := 12
	for i, j := x, y; i <= 14 && j >= 0; {
		Println(i, j)
	}
}

func T72() {
	b := bytes.NewBuffer(nil)
	data := []byte("1234567890")
	n, e := b.Write(data)
	if e != nil {
		Println(e)
	}

	r := b.Next(4)
	Println(n, r, data, b.String(), b.Len())
}

// 模拟一个最小执行时间的阻塞函数
func inc(a int) int {
	res := a + 1                // 虽然我只做了一次简单的 +1 的运算,
	time.Sleep(1 * time.Second) // 但是由于我的机器指令集中没有这条指令,
	// 所以在我执行了 1000000000 条机器指令, 续了 1s 之后, 我才终于得到结果。B)
	return res
}

// 向外部提供的阻塞接口
// 计算 a + b, 注意 a, b 均不能为负
// 如果计算被中断, 则返回 -1
func Add(ctx context.Context, a, b int) int {
	res := 0
	for i := 0; i < a; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			Println("=== a Done")
			return -1
		default:
		}
	}
	for i := 0; i < b; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			Println("=== b Done")
			return -1
		default:
		}
	}
	return res
}

func T73() {
	{
		//使用开放的 API 计算 a+b
		a := 3
		b := 2
		timeout := 2 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		res := Add(ctx, a, b)
		Printf("Compute: %d+%d, result: %d\n", a, b, res)
	}
	{
		// 手动取消
		a := 3
		b := 2
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			cancel() // 在调用处主动取消
		}()
		res := Add(ctx, a, b)
		Printf("Compute: %d+%d, result: %d\n", a, b, res)
	}
}

func T74() {
	s := "site_uv_report"
	n := strings.Index(s, "_")
	s, s1 := s[:n+1], s[n+1:]
	s = strings.Replace(s, "_", "_group_", -1)
	Println(s + s1)
}

type SafeInt struct {
	sync.Mutex
	Num int
}

func T75() {
	count := SafeInt{}
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(i int) {
			count.Lock()
			count.Num += i
			count.Unlock()
			Print(count.Num, i, ", ")
			//Print(count.Num, ", ")
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func T76() {
	data := [...]int{1, 2, 3, 4, 5}
	for i := range data {
		Println(&data[i])
	}

	Println("-------")
	sli := data[:2:2]
	for i := range sli {
		Println(&sli[i])
	}
	Println("-------")

	sli = append(sli, 100)
	for i := range sli {
		Println(&sli[i])
	}

	Println("-------")
	Println(data, sli, cap(sli))
}

func T77() {
LOOP:
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			Println(i)
			break LOOP
		case 4:
			Println(i)
		}
	}

	Println("end")
}

func T78() {
	// You can edit this code!
	// Click here and start typing.
	var str = "dsf士大夫方法啊啊aa"

	Println(strings.Index(str, "方法"))

}

func T79() {
	m := make(map[int]*int)
	a := 1
	b := 2

	m[1] = &a
	m[2] = &b
	for i, v := range m {
		Println(i, v, &(*v), m[i], &a, &b)
	}
}

func T80() {
	Println("Hello World!")
}
