package main

import (
	"fmt"
)

func status(name string) {
	balls := map[string]int{"abcd": 20, "efgh": 0}
	//	ball := balls[name]
	//	fmt.Println(ball)
	ball, exists := balls[name]
	if !exists {
		fmt.Println(name, "out")
	} else if ball < 1 {
		fmt.Println(name, "is ", balls[name], "out")
	} else {
		fmt.Println(name, "won")
	}
}

func main() {
	//var games map[int]string
	//games = make(map[int]string)

	// games := map[int]string{
	// 	100: "a",
	// 	200: "b",
	// 	300: "c",
	// 	400: "d",
	// 	500: "e",
	// 	600: "f",
	// }

	// games[100] = "a"
	// games[200] = "b"
	// games[300] = "c"
	// games[400] = "d"
	// games[500] = "e"
	// games[600] = "f"

	// name, ok := games[101]
	// fmt.Println(name, ok)

	// for _, v := range games {
	// 	fmt.Println(v)
	// }
	// games[500] = "g"
	// delete(games, 400)

	// for k, v := range games {
	// 	fmt.Println(k, v)
	// }

	status("abcd")
	status("bcde")
	status("efgh")

}
