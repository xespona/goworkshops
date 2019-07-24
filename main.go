package main

import (
	"fmt"
	"github.com/xespona/goworkshops/blackjack"
	"strings"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("wooops something went wrong: %s \n", r)
		}
	}()

	fmt.Println("Welcome to Blackjack Simulator!")
	fmt.Println("How many players will be playing? (min: 1, max: 6). The croupier does not count!!")

	numPlayers := getNumPlayers()
	playerNames := getPlayerNames(numPlayers)

	sleepAndSpacesInOutput(1)

	fmt.Printf("Awesome \\m/, current players are: %s \n", strings.Join(playerNames, ", "))
	fmt.Printf("Starting game... \n")

	game := blackjack.New(playerNames, blackjack.NewPokerDeck(), blackjack.DefaultDecksNumber)

	fmt.Printf("carts dealt: %v \n", game.CurrentStatus())
	fmt.Println("Now we will start playing! Remember, press H to Hit and S to Stand")

	var maxScore int
	for _, playerName := range playerNames {
		// the croupier is auto played ;)
		if playerName == blackjack.Croupier {
			continue
		}

		sleepAndSpacesInOutput(1)

		score := playTurn(game, playerName)

		if score > blackjack.MaxScore {
			fmt.Printf("%s, you have %d points! You are bust!\n", playerName, score)
			continue
		}

		if score > maxScore {
			maxScore = score
		}

		if score == blackjack.MaxScore {
			fmt.Printf("That is awesome %s, you have %d \n", playerName, score)
			continue
		}

		fmt.Printf("%s, you stopped at %d points\n", playerName, score)
	}

	if maxScore > 0 {
		fmt.Println("Now its time for the Croupier to play....")
		sleepAndSpacesInOutput(2)

		score := playCroupier(game, maxScore)

		if score > blackjack.MaxScore {
			fmt.Printf("The Croupier stopped at %dpoints, Croupier bust!\n", score)
		} else {
			fmt.Printf("The Croupier stopped at %dpoints\n", score)
		}

		sleepAndSpacesInOutput(2)

		fmt.Printf("Final game cards: %v\n", game.CurrentStatus())
		fmt.Println("... AND THE WINNER IS.. ")

		sleepAndSpacesInOutput(2)

		winner := game.Winner()
		fmt.Println(game.Winner())

		if winner != blackjack.Croupier && winner != blackjack.Draw {
			fmt.Println("yaaaaaaay! Congrats")
		} else {
			fmt.Println("oooooh, better luck next time")
		}
	} else {
		fmt.Println("Everyone busted, the Croupier wins :(")
	}
}

func playCroupier(game blackjack.PlayableBlackJackSim, maxScore int) int {
	score := game.PlayerPoints(blackjack.Croupier)
	for score < maxScore {
		game.Hit(blackjack.Croupier)
		score = game.PlayerPoints(blackjack.Croupier)
		fmt.Printf("current cards: %v, score for %s: %d \n", game.CurrentStatus(), blackjack.Croupier, score)
	}

	if score == maxScore && score < blackjack.MaxScore - 7 {
		game.Hit(blackjack.Croupier)
		score = game.PlayerPoints(blackjack.Croupier)
		fmt.Printf("current cards: %v, score for %s: %d \n", game.CurrentStatus(), blackjack.Croupier, score)
	}
	return score
}

func playTurn(game blackjack.PlayableBlackJackSim, playerName string) int {
	fmt.Printf("%s, It is your turn!\n", playerName)

	score := game.PlayerPoints(playerName)
	for score < blackjack.MaxScore && strings.ToUpper(readMove(playerName)) != "S" {
		game.Hit(playerName)
		score = game.PlayerPoints(playerName)
		fmt.Printf("current cards: %v, score for %s: %d \n", game.CurrentStatus(), playerName, score)
	}
	return score
}

func sleepAndSpacesInOutput(numSeconds int) {
	time.Sleep(time.Duration(numSeconds) * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func getPlayerNames(numPlayers int) []string {
	var playerNames []string
	playerNames = make([]string, numPlayers)
	fmt.Println("That's great! Now, please enter the name for each Player. Remember, the croupier always goes last ;)")
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Enter the name for player %d \n", i+1)
		_, err := fmt.Scanf("%s", &playerNames[i])
		if err != nil {
			panic(err.Error())
		}
	}
	return playerNames
}

func getNumPlayers() int {
	var numPlayers int

	_, err := fmt.Scanf("%d", &numPlayers)

	if err != nil {
		panic(err.Error())
	}

	if numPlayers < blackjack.MinPlayers || numPlayers > blackjack.MaxPlayers {
		panic(fmt.Sprintf("You had one job..... invalid number of Players: %d, try again", numPlayers))
	}

	return numPlayers
}

func readMove(player string) string {
	var move string

	fmt.Println("Hit (H or Anything) or Stand? (S)", player)

	if _, err := fmt.Scanf("%s", &move); err != nil {
		panic(err)
	}

	return move
}
