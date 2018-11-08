package goworker

import (
	"fmt"
)

func main() {
	i := &Imp{}
	f(i)
}

type I interface {
	Get() int
}

func f(p I) {
	fmt.Println(p.Get())
}

type Imp struct {
}

func (i *Imp) Get() int {
	return 122
}

type Imp2 struct {
}

func (i *Imp2) Get() int {

}
