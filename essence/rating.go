package models

//Структура для json с рейтингом игроков
type Rating struct {
	Name           string `json:"name"`
	CorrectAnswers int    `json:"correctAnswers`
}
