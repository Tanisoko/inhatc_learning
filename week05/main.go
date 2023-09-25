package main

import (
	"fmt"
	"strings"
)

func main() {
	//	var now time.Time = time.Now()
	//	var year int = now.Year()
	//	var month string = now.Month().String()
	//	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())

	Hotspurs := "gm ? j madi?"
	replacePlayer := strings.NewReplacer("?", "son")
	Player := replacePlayer.Replace(Hotspurs)
	fmt.Println(Player)

}
