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

func (game *Game) RandomWord() string {
	game.hiddenWord = arrOfWords[rand.Intn(len(arrOfWords))]
	return game.hiddenWord
}

func (game *Game) CheckSymbol(word string) {

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
	}

	game.result = string(runesResult)
}

func (game *Game) StartGame() {
	word := game.RandomWord()

	for i := 0; i < len([]rune(word)); i++ {
		game.result += "_"
	}

	fmt.Println(game.result)

	for {
		fmt.Println("Введите символ")
		fmt.Scan(&game.scanSymbol)
		game.CheckSymbol(word)
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
