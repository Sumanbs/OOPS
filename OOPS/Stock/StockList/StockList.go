//Description - This program is used to add the stock details to the list
//dynamically or to delete the stock dynamiclly.

package main

import "fmt"

//stock - Declaring a structure of stock datatype.

type stock struct {
	Name             string `json:"Name"`
	Number_of_shares int    `json:"Number_of_shares"`
	Share_price      int    `json:"Share_price"`
	next             *stock
}

//LinkedList - A structure.

type LinkedList struct {
	head  *stock
	count int
}

//DeleteStock - This function is used to Delete the particular node(In this program stock).

func (l *LinkedList) DeleteStock(name string) {
	//fmt.Printf("length = %d", l.count)
	if l.count == 0 {
		fmt.Println("Delete not possible, add stock details") //if there are no elements in List
	} else if l.count == 1 { //if Only one element in the list
		l.head = nil
		l.count--
		fmt.Println("Delete successful..")
	} else if l.head.Name == name {
		l.head = l.head.next
		l.count--
		fmt.Println("Delete successful..")
	} else {
		current := l.head
		prev := current
		for current.Name != name {
			prev = current
			current = current.next
		}
		prev.next = current.next
		l.count--
		fmt.Println("Delete successful..")
	}
}

//AddStock - Add new element to the list.

func (l *LinkedList) AddStock(s1 *stock) {
	if l.count == 0 {
		l.head = s1 //No elements in the list
	} else {
		currentPost := l.head
		for currentPost.next != nil {
			currentPost = currentPost.next
		}
		currentPost.next = s1
	}
	l.count++
	fmt.Println("Added successfully")
}

//Display-Display the list Items

func (l *LinkedList) Display() {
	if l.count == 0 {
		fmt.Println("No items in the stock")
	} else {
		fmt.Println("..........The stock report is......")
		currentPost := l.head
		for currentPost != nil {
			fmt.Println("\n-----------------------------------------\n")
			fmt.Printf("\nName : %s", currentPost.Name)
			fmt.Printf("\nShare_Price : %d", currentPost.Share_price)
			fmt.Printf("\nShares : %d", currentPost.Number_of_shares)
			fmt.Println("\n----------------------------------------\n")
			currentPost = currentPost.next
		}
	}
}

//main - In this user will have the options to add,delete or view the stock details.

func main() {
	l := &LinkedList{}
	var flag bool
	var n int
	flag = true
	for flag {
		fmt.Println("..............................")
		fmt.Println("1.Add stocks to the list")
		fmt.Println("2.Delete stocks from the list")
		fmt.Println("3.View the stocks in the list")
		fmt.Println("4.Exit")

		fmt.Print("Enter your choice :")
		fmt.Scanf("%d", &n)
		if n == 1 {
			//get details from user and add it to list
			var Numberofshares, Shareprice int
			var name string
			fmt.Print("Enter the name : ")
			fmt.Scanf("%s", &name)
			fmt.Print("Enter the Number_of_shares : ")
			fmt.Scanf("%d", &Numberofshares)
			fmt.Print("Enter the Share_price : ")
			fmt.Scanf("%d", &Shareprice)
			s := stock{
				Name:             name,
				Number_of_shares: Numberofshares,
				Share_price:      Shareprice,
			}
			l.AddStock(&s) //call AddStock to add data to list.
		}
		if n == 2 {

			// Get string input from user and delete the particular stock from list.

			var name string
			fmt.Print("Enter the name of the stock you want to delete : ")
			fmt.Scanf("%s", &name)
			l.DeleteStock(name)
		}
		if n == 3 {
			l.Display() //Display the list items
		}
		if n == 4 {
			flag = false //exit
		}
	}
	fmt.Printf("First: %v\n", l.head)
}
