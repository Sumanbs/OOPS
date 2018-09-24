<?php
include "AddressBook.php";
class Utility
{
    /**
     * @lastname - Sorts the address using bubble sort.
     */
    public function lastNameSort()
    {
        $data = file_get_contents('AddressBook.json');
        $json_arr = json_decode($data, true);
        for ($i = 0; $i < count($json_arr) - 1; $i++) {
            for ($j = $i + 1; $j < count($json_arr); $j++) {
                if ($json_arr[$i]["LastName"] > $json_arr[$j]["LastName"]) {
                    Utility::swap($i, $j);
                }
            }
        }
    }
    /**
     * @swap - Swap 2 address.
     */
    public function swap($a, $b)
    {
        $data = file_get_contents('AddressBook.json');
        $json_arr = json_decode($data, true);
        $temp = array();
        $temp = $json_arr[$a];
        $json_arr[$a] = $json_arr[$b];
        $json_arr[$b] = $temp;
        file_put_contents('AddressBook.json', json_encode($json_arr));
    }
    /**
     * @display() - This is used to dispaly the address
     */
    public function display()
    {
        $data = file_get_contents('AddressBook.json');
        $json_arr = json_decode($data, true);
        echo "******************Address**********************\n";
        for ($i = 0; $i < count($json_arr); $i++) {
            echo "First Name = " . $json_arr[$i]["FirstName"] . "\n";
            echo "Last Name = " . $json_arr[$i]["LastName"] . "\n";
            echo "Address ";
            echo "State = " . $json_arr[$i]["address"]["State"] . "\n";
            echo "City = " . $json_arr[$i]["address"]["City"] . "\n";
            echo "Zipcode = " . $json_arr[$i]["address"]["Zipcode"] . "\n";
            echo "PH-NO = " . $json_arr[$i]["address"]["PH-NO"] . "\n";
            echo "**************************************************\n";
        }
    }
    /**
     * @Description - Read in the following message: Hello <<name>>, We have your full
     * name as <<full name>> in our system. your contact number is 91-xxxxxxxxxx.
     * Please,let us know in case of any clarification Thank you BridgeLabz 01/01/2016.
     *  Use Regex to replace name, full name, Mobile#, and Date with proper value.
     */
    public function Replace($FirstName, $LastName, $Number, $date)
    {
        $string = "Read in the following message: Hello <<name>>, We have your full name as <<full name>> in our system. your contact number is 91-xxxxxxxxxx. Please,let us know in case of any clarification Thank you BridgeLabz xx-xx-xxxx.";
        echo "\n";
        $pattern = array();

        $pattern[0] = '/<<name>>/';
        $pattern[1] = '/<<full name>>/';
        $pattern[2] = '/xxxxxxxxxx/';
        $pattern[3] = '/xx-xx-xxxx/';

        $replacements[0] = $FirstName;
        $replacements[1] = $FirstName . " " . $LastName;
        $replacements[2] = $Number;
        $replacements[3] = $date;

        $string = preg_replace($pattern, $replacements, $string); //preg_replace is a default function used to replace certain string.
        echo $string;
    }

    /**
     *@getName- Function to validate the name
     */
    public function getName()
    {
        $name = readline();
        if (!preg_match("/[a-zA-Z ]$/", $name) && (preg_match('/[0-9]/', $name))) {
            echo "Invalid name,Enter again\n";
            return Utility::getName(); //call recursively
        } else {
            return $name;
        }
    }
    /**
     * Function to validate the PhoneNumber
     */
    public function getNumber()
    {
        $Number = readline();
        if (!preg_match("/[0-9]{10}$/", $Number)) {
            echo "Invalid Number,Enter Once again\n";
            return Utility::getNumber(); //call recursively
        } else {
            return $Number;
        }
    }
    /**
     * @getZipCode - Function to validate the Zipcode
     */
    public function getZipCode()
    {
        $Number = readline();
        if (!preg_match("/^\d{6}$/", $Number)) {
            echo "Invalid Number,Enter Once again\n";
            return Utility::getZipCode(); //call recursively
        } else {
            return $Number;
        }

    }
    /**
     * @getDate - Function to validate the Date
     */
    public function getDate()
    {
        $date = readline();
        $dat = array();
        $dat = explode("-", $date);
        if (!checkdate($dat[1], $dat[0], $dat[2])) {
            echo "Invalid date,Enter again\n";
            return Utility::getDate(); //call recursively
        } else {
            return $date;
        }
    }
    /**
     * @getShare - Function to validate the Integer
     */
    public function getShare()
    {
        $share = readline();
        if (!is_numeric($share)) {
            echo "invalid entry";
            return Utility::getShare(); //call recursively
        } else {
            return $share;
        }
    }
    /**
     * @getNum - Function to validate the 1 or 2
     */
    public function getNum()
    {
        $num = readline();
        if ($num == 1 || $num == 2) {
            return $num;
        } else {
            echo "invalid entry";
            return Utility::getNum(); //call recursively
        }
    }

}
