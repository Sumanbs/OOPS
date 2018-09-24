<?php
/*
 * Created @ - 18-09-2018
 * This is used to print the FirstName,LastName,State,City,Pincode in the JSON .This program allows you to add address to the JSON file.
 * Delete the address by First name,and sort the address by last name. It also allows you to edit the address.
 */
include "Add_Address.php";

while (1) {
    echo "\n-----------------Please Select the options-------------------\n";
    echo "1.Open the address Book\n";
    echo "2.Add new address to Addressbook\n";
    echo "3.Edit the address Book\n";
    echo "4.Delete the address\n";
    echo "5.Sort the Address\n";
    echo "6.Exit\n";
    echo "\nEnter Your Option : ";
    $ref = new Add_Address();
    $n = $ref->readSelection();
    if ($n == 1) {
        $ref1 = new ChangeAddress();
        $ref1->display(); //call display function
    }
    if ($n == 2) {
        $ref->Get_address(); //call Get_address method
    }
    if ($n == 3) {
        $ref = new ChangeAddress();
        $ref->display();
        $ref->EditAddress(); //to edit the address
    }
    if ($n == 4) {
        $ref = new DeleteAddress();
        $ref->delete(); //to delete the address
    }
    if ($n == 5) {
        $ref = new ChangeAddress();
        $ref->lastNameSort();
        echo "-------------Sorting completed based on LASTNAME------------------";
    }
    if ($n == 6) {
        break;
    }

}
