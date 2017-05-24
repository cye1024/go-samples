package samples

import "fmt"

type T struct {
	name string
	age  int
}

func (t *T) Say() {
	fmt.Printf("Hi, My name is %s and age is %d. \n", t.name, t.age)
	return
}

type In interface {
	Run()
}

type M struct {
	error
	T
	In
	ok bool
}

type Ins struct {
}

func (i *Ins) Run() {
	fmt.Println("hahahah")
}

func Start() {
	m := new(M)
	m.name = "Amy"
	m.age = 18
	m.Say()

	m.In = new(Ins)
	m.Run()
}
