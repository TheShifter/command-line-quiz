package entity

//This is struct for json with question and answers
type Task struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
