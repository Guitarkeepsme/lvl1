// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

func main() {
	rusUser := &RusUser{}
	application := &Application{}
	rusUser.rusUserTypesReq(application)
}



type RusUser struct {
}

func (ru *RusUser) rusUserTypesReq(app RusText) {
	fmt.Println("User types a request in Russian.")
	app.RusText()
}

type RusText interface {
	RusText()
}

type Application struct {
}

func (ts *Application) reqestSend() {
	fmt.Println("A request is sent to the server.")
}

type 

type autoTranslator struct {
	Translator *Application
}

func (at *autoTranslator) requestInEnglish() {
	fmt.Println("A request has been translated and done.")
	at.Translator.reqestSend()
}
