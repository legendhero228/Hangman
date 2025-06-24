package logic

import (
	"fmt"
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
	counterErrors int
	hiddenWord    string
	scanSymbol    string
	result        string
	symbolCorrect bool
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
	game.symbolCorrect = false

	for i, r := range runes {
		if r == scanRunes {
			runesResult[i] = r
			game.symbolCorrect = true
		}
	}

	if !game.symbolCorrect && game.counterErrors < endGame {
		game.counterErrors++
	}

	game.result = string(runesResult)
}

func (game *Game) checkWin() {
	if game.result == game.hiddenWord { // проверка победы
		fmt.Printf("\n\n\nВы выиграли\n\n\n")
	}

	if game.counterErrors == endGame {
		fmt.Println("Вы повешены")
	}

}

func (game *Game) fillResult() {
	for i := 0; i < len([]rune(game.hiddenWord)); i++ {
		game.result += "_"
	}
}

func (game *Game) printResult() {
	if game.symbolCorrect {
		fmt.Println("Вы угадали букву. Вот слово ")
		fmt.Println(game.result)
	} else if !game.symbolCorrect {
		fmt.Println("Вы не угадали букву")
		fmt.Printf("У вас %d ошибок\n", game.counterErrors)
		game.printHangmanStages()
	}

}

func (game *Game) StartGame() {

	game.hiddenWord = game.randomWord()
	fmt.Println("Правильное слово", game.hiddenWord)
	game.fillResult()
	fmt.Println(game.result)

	for {
		fmt.Println("Введите символ")
		fmt.Scan(&game.scanSymbol)
		game.checkSymbol(game.hiddenWord)
		game.checkWin()
		game.printResult()
	}

}
