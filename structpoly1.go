package main

import "fmt"

type MainInterface interface {
	call(string, string) string
}

//------------------------------------
type MainCaller struct {
	MainInterface
}

//------------------------------------
type Normal struct {
}

func (n *Normal)call(x, y string) (str string){
	str = fmt.Sprintf("%v , %v", x, y)
	return
}

//------------------------------------
type Change struct {
}

func (c *Change)call(x, y string) (str string){
	str = fmt.Sprintf("change %v , %v", y, x)
	return
}

//------------------------------------

func main() {
	var m MainCaller

	fmt.Println("----- Sample 1 -----")
	var n = Normal{}

	m.MainInterface = &n
	fmt.Println(m.call("Hello", "World"))

	fmt.Println()

	fmt.Println("----- Sample 2 -----")
	var c = Change{}

	m.MainInterface = &c
	fmt.Println(m.call("Hello", "World"))
}
