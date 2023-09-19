package main

import "fmt"

// Имеющийся интерфейс
type EnglishSpeaker interface {
	SpeakEnglish() string
}

// Структура, реализующая имеющийся интерфейс
type EnglishPerson struct{}

func (e *EnglishPerson) SpeakEnglish() string {
	return "Hello!"
}

// Целевой интерфейс
type RussianSpeaker interface {
	SpeakRussian() string
}

// Адаптер
type Adapter struct {
	englishSpeaker EnglishSpeaker
}

func (a *Adapter) SpeakRussian() string {
	// Вызываем метод SpeakEnglish у имеющегося интерфейса
	englishText := a.englishSpeaker.SpeakEnglish()
	// Производим необходимые преобразования для получения русского текста
	russianText := convertEnglishToRussian(englishText)
	return russianText
}

// Функция для преобразования английского текста в русский
func convertEnglishToRussian(englishText string) string {
	// Здесь может быть реализована логика преобразования текста
	return "Привет!"
}

// Клиентский код
func main() {
	englishPerson := &EnglishPerson{}
	adapter := &Adapter{englishSpeaker: englishPerson}
	russianText := adapter.SpeakRussian()
	fmt.Println(russianText) // Выводит: Привет!
}
