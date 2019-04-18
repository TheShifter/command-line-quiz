package main

import (
	. "command-line-quiz/gameUtil"
	"fmt"
	"github.com/TheShifter/command-line-quiz/gameUtil"
	_ "github.com/TheShifter/command-line-quiz/gameUtil"
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
		gameUtil.s
		tart()
	case 2:
		rating := gameUtil.GetRating()
		fmt.Println(gameUtil.GetTopFive(rating))
	case 3:
		os.Exit(3)
	default:
		fmt.Println("Invalid choice")
	}
}
