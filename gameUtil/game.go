package gameUtil

import (
	"github.com/TheShifter/command-line-quiz/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

const (
	taskFile   = "json/tasks.json"
	ratingFile = "json/rating.json"
)

func getQuestion() (tasks []entity.Task) {
	jsonFile, err := os.Open(taskFile)
	defer jsonFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonVal, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(jsonVal, &tasks)
	if err != nil{
		log.Fatal(err)
	}
	return
}

func calculateQuestion(countCorrect int, countIncorect int) {
	var userAnswer string
	questions := getQuestion()
	shuffle(questions)
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

func Start() {
	var correct int
	var incorect int
	var name string
	timer := time.NewTimer(time.Minute)
	go func() {
		<-timer.C
		go calculateQuestion(&correct, &incorect)
	}()

	if topFive(correct) {
		fmt.Println("Enter your name: ")
		fmt.Fscan(os.Stdin, &name)
		addToRating(name, correct)
	}
	fmt.Printf("Final result:\n"+
		"count of correct answers = %d\n"+"count of incorrect answers = %d", correct, incorect)
}

func GetRating() (ratings []entity.Rating) {
	jsonfile, err := os.Open(ratingFile)
	defer jsonfile.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonVal, _ := ioutil.ReadAll(jsonfile)
	_ = json.Unmarshal(jsonVal, &ratings)
	return
}

func GetTopFive(ratings []entity.Rating) (topFive []entity.Rating) {
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
	rating := GetRating()
	topfive := GetTopFive(rating)
	for _, rating := range topfive {
		if countUserCorrectAnswers >= rating.Correct {
			return true
		}
	}
	return false
}

func addToRating(name string, countCorrectAnsw int) {
	initailRating := GetRating()
	initailRating = append(initailRating, entity.Rating{Name: name, Correct: countCorrectAnsw})
	result, err := json.Marshal(initailRating)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(ratingFile)
	file.WriteString(string(result))
	defer file.Close()
	fmt.Println("You was added to top!!!")
}

func shuffle(tasks []entity.Task)  {
	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})
}
