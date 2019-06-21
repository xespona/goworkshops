package main

import (
	"fmt"
	"github.com/xespona/goworkshops/blackjack"
	"strings"
)

func main() {
	deck := blackjack.NewPockerDeck()

	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Blackjack Simulator!")
	fmt.Println("How many players will be playing? (min: 1, max: 6). The croupier does not count!!")
	var numPlayers int
	_, err := fmt.Scanf("%d", &numPlayers)

	if err != nil {
		fmt.Printf("wooops something went wrong: %s", err.Error())
		return
	}

	//if numPlayers < 1 || numPlayers > 6 {
	//	fmt.Printf("You had one job..... invalid number of Players: %d, try again", numPlayers)
	//	return
	//}

	var playerNames []string
	playerNames = make([]string, numPlayers)

	fmt.Println("That's great! Now, please enter the name for each Player. Remember, the croupier always goes last ;)")
	for i:=0; i< numPlayers; i++ {
		fmt.Printf("Enter the name for player %d \n", i+1)
		_, err = fmt.Scanf("%s", &playerNames[i])
		if err != nil {
			fmt.Printf("wooops something went wrong: %s", err.Error())
			return
		}
	}

	fmt.Printf("Awesome \\m/, current players are: %s \n", strings.Join(playerNames, ", "))
	fmt.Printf("Dealing the first cards for each player... \n");

	game := blackjack.New(playerNames, deck, 1)
	fmt.Printf("%v", game.CurrentStatus())
	fmt.Println("Now we will start playing, yay!! Remember, press H to Hit and S to Stand")

	//var move string
	//for _, err := fmt.Scanf("%s", &move); err != nil && move[1] != 'H' {
	//
	//}

}
