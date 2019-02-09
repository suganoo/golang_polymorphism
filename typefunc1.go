package main

import "fmt"

type MainInterface interface {
	call(string, string) string
}

type SampleFunc func(string, string) string

func (s SampleFunc) call(x, y string) string {
	return s(x, y)
}

func main() {
	var m MainInterface

	fmt.Println("----- Sample 1 -----")
	var sf1 = SampleFunc(func(x, y string) (str string){
		str = fmt.Sprintf("%v , %v", x, y)
		return
	})

	m = sf1
	fmt.Println(m.call("Hello", "World"))

	fmt.Println("----- Sample 2 -----")
	var sf2 = SampleFunc(func(x, y string) (str string){
		str = fmt.Sprintf("Change! %v , %v", y, x)
		return
	})

	m = sf2
	fmt.Println(m.call("Hello", "World"))
}
