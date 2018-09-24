package main

//This program is used to add the date of the transaction using Queue.
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

//struct used for queue
type stock1 struct {
	Name             string    `json:"Name"`
	Number_of_shares int       `json:"Number_of_shares"`
	Share_price      int       `json:"Share_price"`
	Status           string    `json:"Status"`
	Time             time.Time `json:"Time"`
	next             *stock1
}
type LinkedList struct {
	head  *stock1
	count int
}

//This structure is used for stock list and json file
type stock struct {
	Name             string `json:"Name"`
	Number_of_shares int    `json:"Number_of_shares"`
	Share_price      int    `json:"Share_price"`
}

type user struct {
	Username    string `json:"Username"`
	Balance     int    `json:"Balance"`
	Shares      int    `json:"Shares"`
	Share_price int    `json:"Share_price"`
}

func main() {
	l := &LinkedList{}
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
		fmt.Println("7.To view Queue details(Transaction details with time)")
		fmt.Println("8. Exit ")
		fmt.Print("Your choice :")
		fmt.Scanf("%d", &n)

		if n == 1 {
			l.Buy()
		}
		if n == 2 {
			l.Sell()
		}
		if n == 3 {
			CreateUser()
		}
		if n == 4 {
			AppendStock()
		}
		if n == 5 {
			ViewUser()
		}
		if n == 6 {
			ViewStock()
		}
		if n == 7 {
			l.Display()
		}
		if n == 8 {
			flag = false
		}
	}
}

//this is used to add to the stack.
func (l *LinkedList) EnQueue(s1 *stock1) {

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
	//fmt.Printf("count = %d", l.count)
}

//@DeQueue to delete the last transaction of the file.
func (l *LinkedList) DeQueue() {

	if l.count == 0 {
		fmt.Println("No elements in stock")
	} else {
		currentPost := l.head
		prev := currentPost
		for currentPost.next != nil {
			prev = currentPost
			currentPost = currentPost.next
		}
		prev.next = nil
	}
	l.count--
}

//To Display Stock report.
func (l *LinkedList) Display() {
	if l.count == 0 {
		fmt.Println("No transaction happened")
	} else {
		fmt.Println("..........The stock report is......")
		currentPost := l.head
		for currentPost != nil {
			fmt.Println("\n-----------------------------------------\n")
			fmt.Printf("\nName : %s", currentPost.Name)
			fmt.Printf("\nShare_Price : %d", currentPost.Share_price)
			fmt.Printf("\nShares : %d", currentPost.Number_of_shares)
			fmt.Printf("\nStatus : %s \n", currentPost.Status)
			fmt.Println(currentPost.Time)
			fmt.Println("\n----------------------------------------\n")
			currentPost = currentPost.next
		}
	}
}

// This function is used add data to structure.
func (l *LinkedList) AddToQueue(name string, Number_of_shares int, Share_price int, Status string, Time time.Time) {

	s := stock1{
		Name:             name,
		Number_of_shares: Number_of_shares,
		Share_price:      Share_price,
		Status:           Status,
		Time:             Time,
	}
	l.EnQueue(&s)
}

//@GetArrayStocks - This function reads the stock.json file and returns array of stock details.

func GetArrayStocks() []stock {
	//Open the stock.json file
	jsonFile, err := os.Open("stock.json")
	//@err - is to handle the error.

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile) //reads data from the jsonFile till the EOF.
	var s1 []stock
	json.Unmarshal(byteValue, &s1) //Convert structure(string) format.
	return s1

}

//@GetArrayUsers - This function reads the stock.json file and returns array of stock details.

func GetArrayUser() []user {

	jsonFile, err := os.Open("user.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var u1 []user
	json.Unmarshal(byteValue, &u1)
	return u1
}

// @SaveToSock - this function is used to add to the stock.

func SaveToStock(s1 []stock) {
	result, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
	}
	//It is used to write to json file.

	err = ioutil.WriteFile("stock.json", result, 0644)
}

//@SaveToUser - This function updates data to the user.json file.
func SaveToUser(u1 []user) {
	result, err := json.Marshal(u1) //converts string data to a json format.
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("user.json", result, 0644)
}

//This function checks the balance present in user account and asks user to which company share he wants to purchase.
//According to customer purchase user balance and number of stock have been increased ordecreased.
func (l *LinkedList) Buy() {
	fmt.Print("Which company stock do you want to buy :")
	var company_name string
	fmt.Scanf("%s", &company_name)
	check := CheckName(company_name)
	if check == 0 {
		fmt.Println("Invalid name ..Enter again")
		l.Buy()
	} else {
		fmt.Print("Enter the number of share you want purchase :")
		var Number_of_shares int
		fmt.Scanf("%d", &Number_of_shares)
		total_price := 0

		var u1 []user
		u1 = GetArrayUser()
		var s1 []stock
		s1 = GetArrayStocks()
		//calculate total price
		j := 0
		for i := 0; i < len(s1); i++ {
			if s1[i].Name == company_name {
				j = i
				total_price = (s1[i].Share_price) * (Number_of_shares)
				break
			}
		}
		if total_price > u1[0].Balance {
			fmt.Println("Your balance is less,purchase less shares")
		} else {

			u1[0].Shares = u1[0].Shares + Number_of_shares
			u1[0].Balance = u1[0].Balance - total_price
			s1[j].Number_of_shares = s1[j].Number_of_shares - Number_of_shares
			fmt.Println("Purchase successful....")
			fmt.Printf("You bought %d shares", Number_of_shares)
			fmt.Println()

			l.AddToQueue(company_name, Number_of_shares, total_price, "Purchased", time.Now())

		}
		SaveToStock(s1)
		SaveToUser(u1)

	}
}

//@sell-This function is used to sell stocks of the.
//It calculates the stock price and updates the balance and Number of shares in the user account.
func (l *LinkedList) Sell() {
	total_price := 0
	var s1 []stock
	s1 = GetArrayStocks()
	var u1 []user
	u1 = GetArrayUser()
	fmt.Print("How many shares do you wanto sell : ")
	var Number_of_shares int
	fmt.Scanf("%d", &Number_of_shares)

	if u1[0].Shares > 0 && u1[0].Shares < Number_of_shares {
		fmt.Println("You dont have enough stock,Enter less numbers")
	} else {
		fmt.Print("To which company do you want to sell : ")
		var company_name string

		fmt.Scanf("%s", &company_name)
		check := CheckName(company_name)
		if check == 0 {
			fmt.Println("Invalid name ..Enter again")
		} else {
			for i := 0; i < len(s1); i++ {
				if s1[i].Name == company_name {

					total_price = (s1[i].Share_price) * (Number_of_shares)
					s1[i].Number_of_shares = s1[i].Number_of_shares - Number_of_shares
					break
				}
			}
			u1[0].Balance = u1[0].Balance + total_price
			l.AddToQueue(company_name, Number_of_shares, total_price, "Sold", time.Now())
			fmt.Println("Sell is suuccessful and your Balance is updated")
			fmt.Printf("New balance = %d", u1[0].Balance)
			fmt.Println()

		}
		SaveToStock(s1)
		SaveToUser(u1)
	}
}

//@ViewUser - This function Displays user details.
//@GetArrayStock - it returns the array of users.
func CheckName(name string) int {
	jsonFile, err := os.Open("stock.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var s1 []stock
	json.Unmarshal(byteValue, &s1)

	flag := 1
	for i := 0; i < len(s1); i++ {
		if s1[i].Name == name {
			flag = 0
			break
		}
	}
	if flag == 1 {
		return 0
	} else {
		return 1
	}

}

//@ViewUser - This function Displays user details.
//@GetArrayStock - it returns the array of users.
func ViewUser() {
	var u1 []user
	u1 = GetArrayUser()

	fmt.Println("------------------User Details---------------------")
	for i := 0; i < len(u1); i++ {
		fmt.Println("Name:  " + u1[0].Username)
		fmt.Println("Balance: " + strconv.Itoa(u1[0].Balance))
		fmt.Println("Number of shares: " + strconv.Itoa(u1[0].Shares))
		fmt.Println("Share Price: " + strconv.Itoa(u1[0].Share_price))
		fmt.Println("---------------------------------------------------")
	}

}

//This function is used to display the stock details..
//@GetArrayStock - it returns the array of stocks
func ViewStock() {
	var s1 []stock
	s1 = GetArrayStocks()

	fmt.Println("------------------Stock Details---------------------")
	for i := 0; i < len(s1); i++ {
		fmt.Println("Name:  " + s1[i].Name)
		fmt.Println("Number_of_shares: " + strconv.Itoa(s1[i].Number_of_shares))
		fmt.Println("Share_price: " + strconv.Itoa(s1[i].Share_price))
		fmt.Println("---------------------------------------------------")
	}

}

//@CreateUser-This function is used to create the user.json file with user details.
//@u1 is the variable of Structure type - user

func CreateUser() {
	var u1 user
	fmt.Println("Enter the UserName")
	fmt.Scanf("%s", &u1.Username)
	fmt.Println("Enter the Balance")
	fmt.Scanf("%d", &u1.Balance)
	fmt.Println("Enter the Shares")
	fmt.Scanf("%d", &u1.Shares)
	fmt.Println("Enter the Shares price")
	fmt.Scanf("%d", &u1.Share_price)
	b1 := []user{
		{
			Username:    u1.Username,
			Balance:     u1.Balance,
			Shares:      u1.Shares,
			Share_price: u1.Share_price,
		},
	}
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(b1)
	file, err := os.Create("user.json")
	// jsonfile, err := os.Open("file.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	io.Copy(file, buf)
}

//@AppendStock - This function is used Add stock details to the stock.json file.
//It takes stock details from the user.

func AppendStock() {
	var Number_of_shares, Share_price int
	var Name string

	fmt.Println("Enter the Name")
	fmt.Scanf("%s", &Name)
	fmt.Println("Enter the Number_of_shares")
	fmt.Scanf("%d", &Number_of_shares)
	fmt.Println("Enter the Share_price")
	fmt.Scanf("%d", &Share_price)

	jsonFile, err := os.Open("stock.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Stock.json")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var s1 []stock

	// Unmarshall it
	err1 := json.Unmarshal([]byte(byteValue), &s1)
	if err1 != nil {
		fmt.Println(err)
	}
	// add further value into it
	s1 = append(s1, stock{Name: Name, Number_of_shares: Number_of_shares, Share_price: Share_price})

	// now Marshal it
	result, error := json.Marshal(s1)
	if error != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("stock.json", result, 0644)
}
