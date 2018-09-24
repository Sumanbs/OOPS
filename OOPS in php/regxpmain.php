<?php

/**
 * @Description - Read in the following message: Hello <<name>>, We have your full
 * name as <<full name>> in our system. your contact number is 91-xxxxxxxxxx.
 * Please,let us know in case of any clarification Thank you BridgeLabz 01/01/2016.
 *  Use Regex to replace name, full name, Mobile#, and Date with proper value.
 */

include "RegEx.php";
$ref1 = new Utility();
/**
 * Get the Inputs from user.
 */
echo "Enter first name\n";
$FirstName = $ref1->getName();
echo "Enter second name\n";
$LastName = $ref1->getName();

echo "Enter your Mobile number\n";
$Number = $ref1->getNumber();

echo "Enter the date(xx-xx-xxxx)\n";
$date = $ref1->getDate();
/**
 * call the replace function to modify the string.
 */
$ref1->Replace($FirstName, $LastName, $Number, $date);
echo "\n";
