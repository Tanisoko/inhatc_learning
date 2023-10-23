package main

import "fmt"

func double(n *int) {
	*n = *n * 2
}

func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main() {
	var a int = 6
	double(&a)
	fmt.Println(a)

	b := 10
	c := &b
	d := 20
	fmt.Printf("%T\n", c)

	swap(&b, &d)
	fmt.Println(b, d)

}
