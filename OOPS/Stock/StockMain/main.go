/*
* Date 20/09/2018
* @Description : Write a program to read in Stock Names, Number of Share, Share Price.
* Print a Stock Report with total value of each Stock and the total value of Stock.
**/
package main

import (
	"fmt"

	"github.com/StockFile"
)

// main : The user has given the list of choices....

func main() {

	var n int
	var flag bool
	flag = true
	for flag {
		fmt.Println("--------Enter your choice -------------------- ")
		fmt.Println("1.To buy the stock")
		fmt.Println("2.To sell the stock")
		fmt.Println("3.Add user")
		fmt.Println("4.Add stock")
		fmt.Println("5.To view User details")
		fmt.Println("6.To view Stock details")
		fmt.Println("7. Exit ")
		fmt.Print("Your choice :")
		fmt.Scanf("%d", &n)

		if n == 1 {
			StockFile.Buy() //calling buy function
		}
		if n == 2 {
			StockFile.Sell() //calling sell function
		}
		if n == 3 {
			StockFile.CreateUser() //calling createUser function
		}
		if n == 4 {
			StockFile.AppendStock() //calling createUser function
		}
		if n == 5 {
			StockFile.ViewUser() //calling createUser function
		}
		if n == 6 {
			StockFile.ViewStock() //calling createUser function
		}
		if n == 7 {
			flag = false
		}
	}

}
