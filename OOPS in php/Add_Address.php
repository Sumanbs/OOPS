<?php
include "RegEx.php";

class Add_Address extends Utility
{
/**
 * This function reads the number between 1 to 6. 
*/
    public function readSelection()
    {
        $n = readline();
        if ($n >= 1 && $n <= 6) {
            return $n;
        } else {
            echo "Invalid entry,Enter again\n";
            return Add_Address::readSelection();
        }
    }
    /**
    * This function is used the address details from user... 
    */
    public function Get_address()
    {
        echo "Enter the First Name : ";
        $FirstName = Add_Address::getName();

        echo "Enter the Last Name : ";
        $LastName = Add_Address::getName();

        echo "Enter the State : ";
        $state = Add_Address::getName();

        echo "Enter the city : ";
        $city = Add_Address::getName();

        echo "Enter the ZipCode : ";
        $zipcode = Add_Address::getZipCode();

        echo "Enter the phone_Number : ";
        $phone_Number = Add_Address::getNumber();
        /**
        * Pass the address details to the Address1 class.
        */
        $ref = new Address1($FirstName, $LastName, $state, $city, $zipcode, $phone_Number);
        
        $data = file_get_contents('AddressBook.json');//get the contents from AddressBook to $data.
        $json_arr = json_decode($data, true);
        $json_arr[] = $ref;
        file_put_contents('AddressBook.json', json_encode($json_arr));

        $flag = 1;
        while ($flag) {
            echo "\nEnter the options\n";
            echo "1.To view the address book\n";
            echo "2.To add the address address again\n";
            echo "3.Return to main menu\n";
            echo "4.Sort the address book\n";
            echo "Your option : ";
            $n = readline();
            if ($n >= 1 || $n <= 4) {
                switch ($n) {
                    case 1:
                        Add_Address::display();
                        break;
                    case 2:
                        Add_Address::Get_address();
                        break;
                    case 3:$flag = 0;
                        break;
                    case 4:Add_Address::lastNameSort();//call sort function(Sorts based on LastName)
                        break;
                    default:
                        echo "Invalid selection";
                        break;
                }
            }
        }
    }
}
class ChangeAddress extends Utility
{
    /**
    * Reads input in the range between 1-4.
    */
    public function readSelection1()
    {
        $n = readline();
        if ($n >= 1 && $n <= 4) {
            return $n;
        } else {
            echo "Invalid entry,Enter again\n";
            return ChangeAddress::readSelection1();
        }
    }
    /**
    * @EditAddress - This function is used to change the address.
    */
    public function EditAddress()
    {
        $flag = 1;
        echo "Enter the name of person,whose address need to be modified : ";
        $name = ChangeAddress::getName();
        $data = file_get_contents('AddressBook.json');
        $json_arr = json_decode($data, true);
        for ($i = 0; $i < count($json_arr); $i++) {
            if ($json_arr[$i]["FirstName"] == $name) {
                $flag = 0;
                echo "Enter your Choice :\n";
                echo "1.To change state\n";
                echo "2.To Change city\n";
                echo "3.To Change zipcode\n";
                echo "4.To Change phone_Number\n";
                $n = ChangeAddress::readSelection1();
                switch ($n) {
                    /**
                    * To change the state
                    */
                    case 1:
                        echo "Enter the State : ";
                        $state = Utility::getName();
                        $json_arr[$i]["address"]["State"] = $state;
                        file_put_contents('AddressBook.json', json_encode($json_arr));
                        break;
                    case 2:
                    /**
                    * To change the city
                    */
                        echo "Enter the city : ";
                        $City = Utility::getName();
                        $json_arr[$i]["address"]["City"] = $City;
                        file_put_contents('AddressBook.json', json_encode($json_arr));
                        break;
                    case 3:
                    /**
                    * To change the Zipcode
                    */
                        echo "Enter the Zipcode : ";
                        $zipcode = Utility::getZipCode();
                        $json_arr[$i]["address"]["Zipcode"] = $zipcode;
                        file_put_contents('AddressBook.json', json_encode($json_arr));
                        break;

                    case 4:
                    /**
                    * To change the Phone Number
                    */
                        echo "Enter the phone_Number : ";
                        $phone_Number = Utility::getNumber();
                        $json_arr[$i]["address"]["PH-NO"] = $phone_Number;
                        file_put_contents('AddressBook.json', json_encode($json_arr));
                        break;

                    default:
                        echo "Invalid Selection";
                        break;

                }
            }
        }
        /**
        * Below part will execute if the user enters invalid name.
        */
        if ($flag == 1) {
            echo "No such name exists......\n";
            $flag = 1;
            while ($flag) {
                echo "\nEnter the options\n";
                echo "1.To view the address book\n";
                echo "2.To edit the address Book\n";
                echo "3.Return to main menu\n ";
                echo "Your option : ";
                $n = readline();
                if ($n >= 1 || $n <= 3) {
                    switch ($n) {
                        case 1:
                            ChangeAddress::display();
                            break;
                        case 2:
                            ChangeAddress::EditAddress();
                            break;
                        case 3:$flag = 0;
                            break;
                        default:
                            echo "Invalid selection";
                            break;
                    }
                }
            }
        } else {
            $flag = 1;
            while ($flag) {
                echo "\nEnter the options\n";
                echo "1.To view the address book\n";
                echo "2.To sort the address book\n";
                echo "3.Return to main menu\n ";
                echo "Your option : ";
                $n = readline();
                if ($n >= 1 || $n <= 3) {
                    switch ($n) {
                        case 1:
                            ChangeAddress::display();
                            break;

                        case 2:
                            ChangeAddress::lastNameSort();
                            break;

                        case 3:$flag = 0;
                            break;

                        default:
                            echo "\nInvalid Selection\n";
                            break;
                    }
                }

            }

        }
    }
}
class DeleteAddress extends Utility
{
    /**
    * @delete - This function is delete the particular address if the user enters valid name. 
    */
    public function delete()
    {
        $flag = 1;
        $json1_arr = array();
        echo "Enter the FirstName of the person to delete the address:";
        $name = DeleteAddress::getName();
        $data = file_get_contents('AddressBook.json');//get the json file data
        $json_arr = json_decode($data, true);//converts it to array of Objects
        for ($i = 0; $i < count($json_arr); $i++) {
            if ($json_arr[$i]["FirstName"] == $name) {
                $flag = 0;
            } else {
                $json1_arr[] = $json_arr[$i];
            }
        }
        /**
        * If the given name not found below codes will be executed.
        */
        if ($flag == 1) {
            echo "********No such name exists,Enter Proper name***********\n";
            $flag1 = 1;
            while ($flag1) {
                echo "\nEnter the options\n";
                echo "1.To view the address book\n";
                echo "2.To Delete address \n";
                echo "3.Return to main menu\n ";
                echo "Your option : ";
                $n = readline();
                if ($n >= 1 || $n <= 3) {
                    switch ($n) {
                        case 1:
                            DeleteAddress::display();
                            break;
                        case 2:
                            DeleteAddress::delete();
                            break;
                        case 3:$flag1 = 0;
                            break;
                        default:
                            echo "\nInvalid Selection\n";
                            break;
                    }
                }
            }
            ChangeAddress::display();//call display function
            return ChangeAddress::EditAddress();
        } else {
            echo "Deletion is Successful : \n";
            file_put_contents('AddressBook.json', json_encode($json1_arr));
        }
    }
}
