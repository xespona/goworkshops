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

	deck := blackjack.NewPokerDeck()

	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Blackjack Simulator!")
	fmt.Println("How many players will be playing? (min: 1, max: 6). The croupier does not count!!")
	var numPlayers int
	_, err := fmt.Scanf("%d", &numPlayers)

	if err != nil {
		panic(err.Error())
	}

	if numPlayers < blackjack.MinPlayers || numPlayers > blackjack.MaxPlayers {
		panic(fmt.Sprintf("You had one job..... invalid number of Players: %d, try again", numPlayers))
	}

	var playerNames []string
	playerNames = make([]string, numPlayers)

	fmt.Println("That's great! Now, please enter the name for each Player. Remember, the croupier always goes last ;)")
	for i:=0; i< numPlayers; i++ {
		fmt.Printf("Enter the name for player %d \n", i+1)
		_, err = fmt.Scanf("%s", &playerNames[i])
		if err != nil {
			panic(err.Error())
		}
	}

	time.Sleep(1*time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Printf("Awesome \\m/, current players are: %s \n", strings.Join(playerNames, ", "))
	fmt.Printf("Dealing the first cards for each player... \n")

	game := blackjack.New(playerNames, deck, 1)
	fmt.Printf("%v \n", game.CurrentStatus())
	time.Sleep(1*time.Second)
	fmt.Println("Now we will start playing! Remember, press H to Hit and S to Stand")

	//var move string
	//for _, err := fmt.Scanf("%s", &move); err != nil && move[1] != 'H' {
	//
	//}

	var maxScore int
	var score int
	for _, playerName := range playerNames {
		// the croupier is auto played ;)
		if playerName == blackjack.Croupier {
			continue
		}

		time.Sleep(1*time.Second)
		fmt.Print("\033[H\033[2J")
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()

		fmt.Printf("It is your turn %s!\n", playerName)
		score = 0
		for score < blackjack.MaxScore && readMove(playerName) != "S" {
			game.Hit(playerName)
			score = game.PlayerPoints(playerName)
			fmt.Printf("current cards: %v, score for %s: %d \n", game.CurrentStatus(), playerName, score)
		}

		if score > blackjack.MaxScore {
			fmt.Printf("%s, you have %d points! You are bust!\n", playerName, score)
			continue
		} else {
			fmt.Printf("%s, you stopped at %d points\n", playerName, score)
		}

		if score > maxScore {
			maxScore = score
		}
	}

	if maxScore > 0 {
		fmt.Println("Now its time for the Croupier to play....")
		time.Sleep(2*time.Second)
		fmt.Print("\033[H\033[2J")
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()

		score = game.PlayerPoints(blackjack.Croupier)
		for score < maxScore {
			game.Hit(blackjack.Croupier)
			score = game.PlayerPoints(blackjack.Croupier)
			fmt.Printf("current cards: %v, score for %s: %d \n", game.CurrentStatus(), blackjack.Croupier, score)
		}

		if score > blackjack.MaxScore {
			fmt.Printf("The Croupier has more than %dpoints, Croupier bust!\n", blackjack.MaxScore)
		} else {
			fmt.Printf("The Croupier stopped at %dpoints\n", blackjack.MaxScore)
		}

		time.Sleep(2*time.Second)
		fmt.Print("\033[H\033[2J")

		fmt.Printf("Final game cards: %v\n", game.CurrentStatus())
		fmt.Println("AND THE WINNER IS")
		time.Sleep(2*time.Second)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		winner := game.Winner()
		fmt.Println(game.Winner())
		if winner != blackjack.Croupier && winner != blackjack.Draw {
			fmt.Println("yaaaaaaay! Congrats")
		} else {
			fmt.Println("oooooh, better luck next time")
		}
	} else {
		fmt.Println("Everyone busted, the Croupier wins. You should learn to play :)")
	}


}

func readMove(player string) string {
	var move string

	fmt.Println("Hit (H) or Stand? (S)", player)
	if _, err := fmt.Scanf("%s", &move); err != nil {
		panic(err)
	}

	return move
}
