package Stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type stock struct {
	Name             string `json:"Name"`
	Number_of_shares int    `json:"Number_of_shares"`
	Share_price      int    `json:"Share_price"`
}

type user struct {
	Username string `json:"Username"`
	Balance  int    `json:"Balance"`
	Shares   int    `json:"Shares"`
}

/*
*@GetArrayStocks - This function reads the stock.json file and returns array of stock details.
 */
func GetArrayStocks() []stock {
	//Open the stock.json file
	jsonFile, err := os.Open("stock.json")
	/*
	*@err - is to handle the error.
	 */
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile) //reads data from the jsonFile till the EOF.
	var s1 []stock
	json.Unmarshal(byteValue, &s1) //Convert structure(string) format.
	return s1

}

/*
*@GetArrayUsers - This function reads the stock.json file and returns array of stock details.
 */
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

/*
* @SaveToSock - this function is used to add to the stock.
 */
func SaveToStock(s1 []stock) {
	result, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
	}
	//It is used to write to json file.
	err = ioutil.WriteFile("stock.json", result, 0644)
}

/*
*@SaveToUser - This function updates data to the user.json file.
 */
func SaveToUser(u1 []user) {
	result, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("users.json", result, 0644)
}

/*
*This function checks the balance present in user account and asks user to which company share he wants to purchase.
*According to customer purchase user balance and number of stock have been increased ordecreased.
 */
func Buy() {
	fmt.Println("Which company stock do you want to buy ")
	var company_name string
	fmt.Scanf("%s", &company_name)
	check := checkname(company_name)
	if check == 0 {
		fmt.Println("Invalid name ..Enter again")
		Buy()
	} else {
		fmt.Println("Enter the number of share you want purchase")
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
				fmt.Println("Executing 169")
				j = i
				total_price = (s1[i].Share_price) * (s1[i].Number_of_shares)
				break
			}
		}
		if total_price > u1[0].Balance {
			fmt.Println("Your balance is less,purchase less shares")
		} else {
			u1[0].Shares = u1[0].Shares + Number_of_shares
			u1[0].Balance = u1[0].Balance - total_price
			s1[j].Number_of_shares = s1[j].Number_of_shares - Number_of_shares

		}
		SaveToStock(s1)
		SaveToUser(u1)
	}
}
func checkname(name string) int {
	jsonFile, err := os.Open("stock.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var s1 []stock
	json.Unmarshal(byteValue, &s1)
	fmt.Println(s1)
	flag := 1
	for i := 0; i < len(s1); i++ {
		if s1[i].Name == name {
			fmt.Printf("Available stocks are : %d", strconv.Itoa(s1[i].Number_of_shares))
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
func Read() {
	// Open our jsonFile
	jsonFile, err := os.Open("file.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened file.json")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var s1 []stock

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &s1)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	for i := 0; i < len(s1); i++ {
		fmt.Println("Name:  " + s1[i].Name)
		fmt.Println("Number of stocks: " + strconv.Itoa(s1[i].Number_of_shares))
		fmt.Println("Number of shares: " + strconv.Itoa(s1[i].Share_price))
	}

	for i := 0; i < len(s1); i++ {

		s1[i].Name = "asdfas"
		fmt.Println("Name:  " + s1[i].Name)
	}
	result, error := json.Marshal(s1)
	if error != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("file.json", result, 0644)
	defer jsonFile.Close()
}
func Appenduser() {
	var Balanc, Share int
	var username string

	fmt.Println("Enter the UserName")
	fmt.Scanf("%s", &username)
	fmt.Println("Enter the Balance")
	fmt.Scanf("%d", &Balanc)
	fmt.Println("Enter the Shares")
	fmt.Scanf("%d", &Share)

	jsonFile, err := os.Open("user.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened User.json")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var u1 []user

	// Unmarshall it
	err1 := json.Unmarshal([]byte(byteValue), &u1)
	if err1 != nil {
		fmt.Println(err)
	}
	// add further value into it
	u1 = append(u1, user{Username: username, Balance: Balanc, Shares: Share})

	// now Marshal it
	result, error := json.Marshal(u1)
	if error != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("user.json", result, 0644)
}
