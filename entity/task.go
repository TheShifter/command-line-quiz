package entity

// Структура для json с вопросами и ответами
type Task struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
