package gameUtil

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TheShifter/command-line-quiz/entity"
	"io/ioutil"
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
		errors.New("File not found" + err.Error())
	}
	jsonVal, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(jsonVal, &tasks)
	if err != nil{
		errors.New("Unmarshal failed" + err.Error())
	}
	return
}

func calculate() (int, int){
	var correct int
	var incorect int
	var userAnswer string
	timer := time.NewTimer(time.Minute)
	questions := getQuestion()
	shuffle(questions)
	questionLoop:
	for _, Task := range questions{
		fmt.Println(Task.Question)
		answerCh := make(chan string)
		go func(){
			fmt.Fscan(os.Stdin, &userAnswer)
			answerCh <- userAnswer
		}()
		select {
		case <- timer.C:
			fmt.Println("time is over")
			break questionLoop
			case userAnswer := <- answerCh:
				if Task.Answer == userAnswer{
					correct++
				}else{
					incorect++
				}
		}
	}
	return correct, incorect
}

func Start() {
	var name string
	correct, incorect := calculate()
	if topFive(correct) {
		fmt.Println("Enter your name: ")
		fmt.Fscan(os.Stdin, &name)
		addToRating(name, correct)
	}
	fmt.Printf("Final result:\n"+"count of correct answers = %d\n"+"count of incorrect answers = %d", correct, incorect)

}

func GetRating() (ratings []entity.Rating) {
	jsonfile, err := os.Open(ratingFile)
	defer jsonfile.Close()
	if err != nil {
		errors.New("File not found" + err.Error())
	}
	jsonVal, _ := ioutil.ReadAll(jsonfile)
	err = json.Unmarshal(jsonVal, &ratings)
	if err != nil{
		errors.New("Unmarshal failed" + err.Error())
	}
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

func addToRating(name string, correct int) {
	initailRating := GetRating()
	initailRating = append(initailRating, entity.Rating{Name: name, Correct: correct})
	result, err := json.Marshal(initailRating)
	if err != nil {
		errors.New("Marshal was failed" + err.Error())
	}
	file, err := os.Create(ratingFile)
	file.WriteString(string(result))
	defer file.Close()
	fmt.Println("You was added to top!!!")
}

func shuffle(tasks []entity.Task){
	r := rand.New((rand.NewSource(time.Now().Unix())))
	for n := len(tasks); n>0; n--{
		randIndex := r.Intn(n)
		tasks[n-1], tasks[randIndex] = tasks[randIndex], tasks[n-1]
	}
}
