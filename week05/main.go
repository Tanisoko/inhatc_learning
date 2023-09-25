package main

import (
	"bufio"
	"fmt"
	"log"
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
	inputScore, err := reader.ReadString('\n')
	//	inputScore, _ := reader.ReadString('\n') // 1 variable but reader.ReadString returns 2 values
	if err != nil {
		log.Fatal(err)
	}

	if inputScore >= 90 { //	mismatched types string and untyped int
		grade := "A grade!"
	} else {
		grade := "BCDE grade~"
	}
	fmt.Println(inputScore)

}
