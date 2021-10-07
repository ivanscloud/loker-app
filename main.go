package main

import (
	"errors"
	"fmt"
	"loker-app/controllers"
	"loker-app/helpers"
	"strings"
)

func main() {
	// exit message
	defer fmt.Println("Terima kasih telah menggunakan Loker-App!")
	// welcome message
	fmt.Println("Welcome to Loker-App. Masukkan perintah Init!")
	var userInput []string
	for {
		userInput = helpers.GetInput()
		if !running(userInput) {
			break
		}
	}
}

func running(userInput []string) bool {

	command, arg1, arg2 := helpers.CheckCommand(userInput)
	command = strings.ToLower(command)
	var err error
	switch command {

	case "init":
		controllers.Init(arg1)
	case "status":
		controllers.Status()
	case "input":
		controllers.Input(arg1, arg2)
	case "leave":
		controllers.Leave(arg1)
	case "find":
		controllers.Find(arg1)
	case "search":
		controllers.Search(arg1)
	case "exit":
		return false
	case "":
		err = errors.New("perintah tidak boleh kosong")
		helpers.ShowResult("", err)
	default:
		err = errors.New("perintah yang Anda masukkan salah")
		helpers.ShowResult("", err)
	}
	return true
}
