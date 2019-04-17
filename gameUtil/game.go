package gameUtil

import (
	. "command-line-quiz/essence"
	"encoding/json"
	"fmt"
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

func calculateQuestion(countCorrectAnsw *int, countIncorectAnsv *int) {
	var userAnswer string
	questions := getQuestion()
	for _, Task := range questions {
		fmt.Println(Task.Question)
		fmt.Fscan(os.Stdin, &userAnswer)
		if Task.Answer == userAnswer {
			*countCorrectAnsw++
		} else {
			*countIncorectAnsv++
		}
	}
}

func Start() {
	var countCorrectAnsw int
	var countIncorectAnsv int
	var name string
	go calculateQuestion(&countCorrectAnsw, &countIncorectAnsv)
	time.Sleep(time.Minute)
	if TopFive(countCorrectAnsw) {
		fmt.Println("Enter your name: ")
		fmt.Fscan(os.Stdin, &name)
		addToRating(name, countCorrectAnsw)
	}
	fmt.Printf("Final result:\n"+
		"count of correct answers = %d\n"+"count of incorrect answers = %d", countCorrectAnsw, countIncorectAnsv)
}

func GetRating() (ratings []Rating) {
	jsonfile, err := os.Open(ratingFile)
	defer jsonfile.Close()
	if err != nil {
		panic(err)
	}
	jsonVal, _ := ioutil.ReadAll(jsonfile)
	_ = json.Unmarshal(jsonVal, &ratings)
	return
}

func GetTopFive(ratings []Rating) (topFive []Rating) {
	sort.Slice(ratings, func(i, j int) bool {
		return ratings[i].CorrectAnswers > ratings[j].CorrectAnswers
	})
	if len(ratings) >= 5 {
		topFive = ratings[0:5]
		return
	} else {
		return ratings
	}
}

func TopFive(countUserCorrectAnswers int) bool {
	rating := GetRating()
	topfive := GetTopFive(rating)
	for _, rating := range topfive {
		if countUserCorrectAnswers >= rating.CorrectAnswers {
			return true
		}
	}
	return false
}

func addToRating(name string, countCorrectAnsw int) {
	initailRating := GetRating()
	initailRating = append(initailRating, Rating{Name: name, CorrectAnswers: countCorrectAnsw})
	result, err := json.Marshal(initailRating)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(ratingFile)
	file.WriteString(string(result))
	defer file.Close()
	fmt.Println("You was added to top!!!")
}
