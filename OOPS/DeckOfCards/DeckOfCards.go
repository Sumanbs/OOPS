package main

//@discription  : DeckOfCards
//suit : array of cards type
//rank : array of cards numbers
//deck  : array of full cards

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	PlayerCards()
}
func PlayerCards() {
	suit := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	rank := [13]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	deck := [52]string{}
	l := 0
	for i := 0; i < len(suit); i++ {
		for j := 0; j < len(rank); j++ {
			//creating card using array
			deck[l] = suit[i] + " of " + rank[j]
			l++
		}
	}

	//players : 2d array
	//size of inner array : no of players
	//size of outer array : no of cards to be distrubuted

	players := [9][4]string{}
	//store in temp deck to shuffel
	tempdeck := deck

	for i := 0; i < len(deck); i++ {
		//random function
		rand.Seed(time.Now().UTC().UnixNano())
		r := rand.Intn(51 - 0)
		var temp string
		//shuffle cards
		temp = tempdeck[r]
		tempdeck[r] = tempdeck[i]
		tempdeck[i] = temp
	}
	num := 0
	fmt.Println("*************************   PLAYERS AND THEIR CARDS DETAILS   ***********************************")
	for i := 0; i < 9; i++ {
		for j := 0; j < 4; j++ {
			//assigning to each players
			players[i][j] = tempdeck[num]
			num++
		}
	}
	temp := [4][9]string{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 4; j++ {
			temp[j][i] = players[i][j]
		}
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("***************Player %d**************\n", (i + 1))

		for j := 0; j < 9; j++ {
			fmt.Println(temp[i][j])
		}

		fmt.Println()
	}
}

// func randInt(min int, max int) int {
// 	//return random number btw min and max
// 	return min + rand.Intn(max-min)
// }
