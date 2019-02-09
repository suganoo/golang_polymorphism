package main

import "fmt"

type MainInterface interface {
	call(string, string) string
}

type SampleFunc func(string, string) string

func (s SampleFunc) call(x, y string) string {
	return s(x, y)
}

func normalFunc(x, y string) (str string){
	str = fmt.Sprintf("%v , %v", x, y)
	return
}
func changeFunc(x, y string) (str string){
	str = fmt.Sprintf("change %v , %v", y, x)
	return
}
func main() {
	var m MainInterface

	fmt.Println("----- Sample 1 -----")
	var sf1 = SampleFunc(normalFunc)

	m = sf1
	fmt.Println(m.call("Hello", "World"))

	fmt.Println("----- Sample 2 -----")
	var sf2 = SampleFunc(changeFunc)

	m = sf2
	fmt.Println(m.call("Hello", "World"))
}
