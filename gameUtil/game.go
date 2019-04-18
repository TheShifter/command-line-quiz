package gameUtil

import (
	. "command-line-quiz/essence"
	"encoding/json"
	"fmt"
	"github.com/TheShifter/command-line-quiz/essence"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

const (
	taskFile   = "json/tasks.json"
	ratingFile = "json/rating.json"
)

func getQuestion() (questions []Task) {
	jsonFile, err := os.Open(taskFile)
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}
	jsonVal, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(jsonVal, &questions)
	return
}

func calculateQuestion(countCorrect *int, countIncorect *int) {
	var userAnswer string
	questions := getQuestion()
	for _, Task := range questions {
		fmt.Println(Task.Question)
		fmt.Fscan(os.Stdin, &userAnswer)
		if Task.Answer == userAnswer {
			*countCorrect++
		} else {
			*countIncorect++
		}
	}
}

func start() {
	var countCorrect int
	var countIncorect int
	var name string
	go calculateQuestion(&countCorrect, &countIncorect)
	time.Sleep(time.Minute)
	if topFive(countCorrect) {
		fmt.Println("Enter your name: ")
		fmt.Fscan(os.Stdin, &name)
		addToRating(name, countCorrect)
	}
	fmt.Printf("Final result:\n"+
		"count of correct answers = %d\n"+"count of incorrect answers = %d", countCorrect, countIncorect)
}

func getRating() (ratings []Rating) {
	jsonfile, err := os.Open(ratingFile)
	defer jsonfile.Close()
	if err != nil {
		panic(err)
	}
	jsonVal, _ := ioutil.ReadAll(jsonfile)
	_ = json.Unmarshal(jsonVal, &ratings)
	return
}

func getTopFive(ratings []Rating) (topFive []Rating) {
	sort.Slice(ratings, func(i, j int) bool {
		return ratings[i].Correct > ratings[j].Correct
	})
	if len(ratings) >= 5 {
		topFive = ratings[0:5]
		return
	} else {
		return ratings
	}
}

func topFive(countUserCorrectAnswers int) bool {
	rating := getRating()
	topfive := getTopFive(rating)
	for _, rating := range topfive {
		if countUserCorrectAnswers >= rating.Correct {
			return true
		}
	}
	return false
}

func addToRating(name string, countCorrectAnsw int) {
	initailRating := getRating()
	initailRating = append(initailRating, Rating{Name: name, Correct: countCorrectAnsw})
	result, err := json.Marshal(initailRating)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(ratingFile)
	file.WriteString(string(result))
	defer file.Close()
	fmt.Println("You was added to top!!!")
}
