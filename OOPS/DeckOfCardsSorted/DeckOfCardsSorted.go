package main

//@discription  : DeckOfCards
//suit : array of cards type
//rank : array of cards numbers
//deck  : array of full cards

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//structure of node type having 2 fields
type node struct {
	data string
	next *node
}
type LinkedList struct {
	head  *node
	count int
}

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
	players := [9][4]string{} //2D array of players
	//store in temp deck to shuffel
	tempdeck := deck
	//for shuffling the cards
	for i := 0; i < len(deck); i++ {
		rand.Seed(time.Now().UTC().UnixNano()) //random function
		r := rand.Intn(51 - 0)
		var temp string
		//shuffle cards
		temp = tempdeck[r]
		tempdeck[r] = tempdeck[i]
		tempdeck[i] = temp
	}
	num := 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 4; j++ {
			//assigning to each players
			players[i][j] = tempdeck[num]
			num++
		}
	}
	temp := [4][9]string{}
	temp1 := [4][9]string{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 4; j++ {
			temp[j][i] = players[i][j]
		}
	}
	t := &LinkedList{}
	for i := 0; i < 4; i++ {
		a := 0
		for k := 0; k < 13; k++ {
			for j := 0; j < 9; j++ {
				s := strings.Split(temp[i][j], " ")
				if rank[k] == s[2] {
					temp1[i][a] = temp[i][j]
					n := node{
						data: temp1[i][a],
					}
					t.EnQueue(&n)
					a++
				}
			}
		}
		fmt.Println()
	}
	for i := 0; i < 4; i++ {
		fmt.Printf("******Player %d************\n", (i + 1))
		for j := 0; j < 9; j++ {
			fmt.Println(t.DeQueue())
		}
		fmt.Println()
	}
}
func (l *LinkedList) EnQueue(s1 *node) {
	if l.count == 0 {
		l.head = s1

	} else {
		currentPost := l.head
		for currentPost.next != nil {
			currentPost = currentPost.next
		}
		currentPost.next = s1

	}
	l.count++
}
func (l *LinkedList) DeQueue() string {
	data := l.head.data
	l.count--
	l.head = l.head.next
	return data
}
