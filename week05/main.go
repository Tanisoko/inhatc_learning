package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//	var now time.Time = time.Now()
	//	var year int = now.Year()
	//	var month string = now.Month().String()
	//	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())

	//	Hotspurs := "gm ? j madi?"
	//	replacePlayer := strings.NewReplacer("?", "son")
	//	Player := replacePlayer.Replace(Hotspurs)
	//	fmt.Println(Player)

	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') // 1 variable but reader.ReadString returns 2 values
	fmt.Println(inputScore, err)

}
