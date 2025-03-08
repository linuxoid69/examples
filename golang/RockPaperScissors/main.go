// build: GOOS=linux GOARCH=amd64 go build -o RockPaperScissors main.go
// build: GOOS=windows GOARCH=amd64 go build -o RockPaperScissors.exe main.go
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

const (
	SELECT_THINGS   = "Выбирете вариант: 1 - Камень, 2 - Ножницы, 3 - Бумага"
	THING_NOT_FOUND = "Такой вариант отсутствует"
	WIN_PC          = "Выйграл Компьютер"
	WIN_USER        = "Выйграли Вы"
	WIN_TIE         = "Ничья"
)

func getDigit(d string) (int, error) {
	digit, err := strconv.Atoi(d)
	if err != nil {
		return 0, err
	}

	return digit, nil
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	cmd.Run()
}

func game(things []string) {
	for {
		var input string

		fmt.Println(SELECT_THINGS)

		fmt.Scanln(&input)

		userDigit, err := getDigit(input)
		if err != nil || userDigit > len(things) {
			clearScreen()

			continue
		}

		pcDigit, _ := rand.Int(rand.Reader, big.NewInt(int64(len(things))))

		var win string

		switch {
		case (userDigit-1 < int(pcDigit.Int64()) && int(pcDigit.Int64()) != 2) ||
			(userDigit-1 == 2 && int(pcDigit.Int64()) == 0):
			win = WIN_USER
		case userDigit-1 == int(pcDigit.Int64()):
			win = WIN_TIE
		default:
			win = WIN_PC
		}

		fmt.Printf("%d, %d\n", userDigit-1, int(pcDigit.Int64()))
		fmt.Printf("Вы выбрали: %12s\nКомпьютер выбрал: %s\n%s\n", things[userDigit-1], things[pcDigit.Int64()], win)
		fmt.Println("Играем еще? - y,n")
		fmt.Scanln(&input)

		if input == "y" {
			clearScreen()

			continue
		} else {
			os.Exit(0)
		}
	}
}

func main() {
	things := []string{"Камень", "Ножницы", "Бумага"}

	clearScreen()
	game(things)
}
