package main

import (
	"fmt"
	"sort"
)

func main() {
	highScores := make([]int, 4) // creating a slice using "make"
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777  ---> this will give an error
	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores) // [234 945 465 867 555 666 321

	highScores = append(highScores, 666)
	fmt.Println(highScores) // 234 945 465 867 555 666 321 666]

	sort.Ints(highScores)
	fmt.Println(highScores) // [234 321 465 555 666 666 867 945]

}
