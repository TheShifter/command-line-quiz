package main

import (
	"errors"
	"fmt"
	"github.com/TheShifter/command-line-quiz/gameUtil"
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
		errors.New("nothing entered" + err.Error())
	}
	switch choise {
	case 1:
		gameUtil.Start()
	case 2:
		rating := gameUtil.GetRating()
		fmt.Println(gameUtil.GetTopFive(rating))
	case 3:
		exit()
	default:
		fmt.Println("Invalid choice")
	}
}

func exit(){
	fmt.Println("You are out of the game ")
	os.Exit(0)
}