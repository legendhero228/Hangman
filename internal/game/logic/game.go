package logic

import (
	"fmt"
	"hangman/internal/game"
	"math/rand"
)

var (
	arrOfWords = []string{
		"арбуз",
		"грузовик",
		"ежевика",
		"пингвин",
		"лампа",
		"картина",
		"компас",
		"корабль",
		"сумерки",
		"ветер",
	}

	hangmanStages = []string{
		`
  +---+
  |   |
      |
      |
      |
      |
=========
`,
		`
  +---+
  |   |
  O   |
      |
      |
      |
=========
`,
		`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========
`,
		`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========
`,
		`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
`,
		`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========
`,
		`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
`,
	}
)

const endGame = 6 // Количество допустимых ошибок

type Game struct {
	game.Word
	game.Player
	counter       int
	counterErrors int
	hiddenWord    string
	scanSymbol    string
	result        string
}

func (game *Game) randomWord() string {
	game.hiddenWord = arrOfWords[rand.Intn(len(arrOfWords))]
	return game.hiddenWord
}

func (game *Game) printHangmanStages() {
	fmt.Println("Hangman Stages:", hangmanStages[game.counterErrors])

}

func (game *Game) checkSymbol(word string) {

	runes := []rune(word)
	scanRunes := []rune(game.scanSymbol)[0]
	runesResult := []rune(game.result)
	temp := false

	for i, r := range runes {
		if r == scanRunes {
			runesResult[i] = r
			fmt.Println("Вы угадали букву: ", game.scanSymbol)
			fmt.Println("Вот угаданные буквы:", string(runesResult))
			game.counter++
			temp = true
		}
	}

	if !temp {
		game.counterErrors++
		fmt.Println("Каунтер ошибок:", game.counterErrors) // todo: Remove
		game.printHangmanStages()
	}

	game.result = string(runesResult)
}

func (game *Game) StartGame() {

	word := game.randomWord()

	for i := 0; i < len([]rune(word)); i++ {
		game.result += "_"
	}

	fmt.Println(game.result)

	for {
		fmt.Println("Введите символ")
		fmt.Scan(&game.scanSymbol)
		game.checkSymbol(word)

		if game.counter == len([]rune(word)) { // проверка победы
			fmt.Println("Вы выиграли")
			return
		}

		if game.counterErrors == endGame {
			fmt.Println("Вы повешены")
			return
		}
	}

}
