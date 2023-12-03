package main

import (
	"fmt"
)

func main() {
	var balls map[string]int
	bells := make(map[string]int)

	fmt.Printf("%#v\n", balls) //	맵 변수의 제로 값은 nil
	fmt.Printf("%#v\n", bells) //	키/값을 추가하려면 make 사용 or 맵 리터럴 사용
}
