package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() []string {
	reader := bufio.NewReader(os.Stdin)
	rawUserInput, _ := reader.ReadString('\n')
	userInput := strings.Fields(rawUserInput)

	return userInput
}

func CheckCommand(userInput []string) (command string, arg1 string, arg2 string) {
	// mendeteksi perintah / command

	if len(userInput) == 0 {
		return
	}

	command = userInput[0]

	if len(userInput) > 2 {
		arg1, arg2 = userInput[1], userInput[2]
	} else if len(userInput) > 1 {
		arg1 = userInput[1]
	}
	return
}

func ShowResult(res string, err error) {
	// menampilkan result / error
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
