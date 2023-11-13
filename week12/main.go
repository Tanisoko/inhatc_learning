package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := []int{0, 0, 0, 0, 0}

	// s[2] = 91
	// s[4] = 99

	// copyS := s[1:4]

	// for _, value := range copyS {
	// 	fmt.Println(value)
	// }

	// test := [3]string{"inha", "go", "student"}
	// testS := test[:3]
	// testS2 := test[1:]
	// testS2[0] = "python"
	// fmt.Println(testS, len(testS))
	// fmt.Println(testS2, len(testS2))
	// fmt.Println(testS2)

	// a := []string{"a", "b", "c", "d"}
	// as := a[:2]
	// as[1] = "x"
	// fmt.Println(as)

	a := make([]string, 4, 5)
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	as[1] = "x"
	// c := append(a, "y")
	c := append(a, "y", "S")

	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))

	fmt.Println("%x %x %x\n", &a[0], &as[0], &c[0])

}
