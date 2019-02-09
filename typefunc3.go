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
	str = fmt.Sprintf("Change %v , %v", y, x)
	return
}

//func funcProducer(i int) (m MainInterface) {
func funcProducer(i int) (m SampleFunc) {
	switch i {
	case 1:
		m = SampleFunc(normalFunc)
	case 2:
		m = SampleFunc(changeFunc)
	default:
		m = nil
	}
	return
}
func main() {
	var m MainInterface

	fmt.Println("----- Sample 1 -----")
	var sf1 = funcProducer(1)

	m = sf1
	fmt.Println(m.call("Hello", "World"))

	fmt.Println("----- Sample 2 -----")
	var sf2 = funcProducer(2)

	m = sf2
	fmt.Println(m.call("Hello", "World"))
}
