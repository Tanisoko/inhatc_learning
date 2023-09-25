package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	inputScoreString = strings.TrimSpace(inputScoreString)
	inputScore, err := strconv.ParseFloat(inputScoreString, 32)

	var grade string

	if inputScore >= 90 {
		grade = "A grade!"
	} else {
		grade = "BCDF grade~"
	}
	fmt.Println("You got", grade)

}
