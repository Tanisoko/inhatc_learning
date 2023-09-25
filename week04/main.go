package main

import (
	"fmt"
	"reflect"
)

func main() {

	//	var a int
	//	a = 9
	//	var a int = 9
	a := 9
	b := 2.71
	c := 'X'
	d, e := 4.10, 3.99
	f := "문자열"

	fmt.Println(a, reflect.TypeOf(a), b, reflect.TypeOf(b), c, reflect.TypeOf(c), d, reflect.TypeOf(d), e, reflect.TypeOf(e), f, reflect.TypeOf(f))

	var g int
	var h float32
	var i bool
	var j string
	fmt.Println(g, h, i, j)
	fmt.Println("%d%f%d%s", g, h, i, j)

	// fmt.Println(reflect.TypeOf('X'))
	// fmt.Println(reflect.TypeOf(99))
	// fmt.Println(reflect.TypeOf(2.7))
	// fmt.Println(reflect.TypeOf(false))
	// fmt.Println(reflect.TypeOf("Go!"))
	// fmt.Println('A', 'a', '0', '김', '인', '하')
	// fmt.Println(math.Ceil(3.17))
	//	fmt.Println(math.Floor(3.17))
	//	fmt.Println(strings.Title("오픈소스 프로그래밍~"))
	//	fmt.Println(strings.Title("open source programming~\n\"go!\""))
}
