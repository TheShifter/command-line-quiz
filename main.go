package main

import (
	. "command-line-quiz/gameUtil"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Make your choice:\n" +
		"1. Start game\n" +
		"2. Get rating\n" +
		"3. Exit\n")
	var choise int
	_, err := fmt.Fscan(os.Stdin, &choise)
	if err != nil {
		panic(err)
	}
	switch choise {
	case 1:
		Start()
	case 2:
		rating := GetRating()
		fmt.Println(GetTopFive(rating))
	case 3:
		os.Exit(3)
	default:
		fmt.Println("Invalid choice")
	}
}
