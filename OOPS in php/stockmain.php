<?php
/**
 * Desc -> Write a program to read in Stock Names, Number of Share, Share Price.
 *Print a Stock Report with total value of each Stock and the total value of Stock.
 */
include "stock.php";
include "RegEx.php";
$ref = new Utility();
$num = 1;
$i = 0;
while ($num == 1) {
    /**
     * Take input from user.
     */
    echo "\nEnter 1 to add data to stock report,2 to see the report :";
    echo "Enter your choice :";
    $num = $ref->getNum();
    if ($num == 1) {
        echo "Enter the name :";
        $name = $ref->getName();

        echo "Enter the Number of Share :";
        $share_num = $ref->getShare();

        echo "Enter the Share price :";
        $share_price = $ref->getShare();

        $ref1 = new Stock($name, $share_num, $share_price);
        $ref1->add_to_json();
    }
}
/**
 * Get the data from json file and decode it to array.
 */
$data = file_get_contents('demo.json');
$total = 0;
$json_arr = json_decode($data,true);
/**
 * Display the result.
 */
echo "--------------Stock report----------------------------------------";
for ($i=0; $i <count($json_arr) ; $i++)
{ 
    echo "\nName = " . $json_arr[$i]['name']."\n";
    echo "stock_price = " . $json_arr[$i]['stock_price']."\n";
    echo "number_of_shares = " . $json_arr[$i]['number_of_shares']."\n";
    echo "total_price = " . ($json_arr[$i]['stock_price']*$json_arr[$i]['number_of_shares'])."\n";
    echo "----------------------------------------------------------";
    $total = $total + ($json_arr[$i]['stock_price']*$json_arr[$i]['number_of_shares']);
}
//print_r($json_arr);
echo "\n\nThe total Stock Price is :";

echo $total."\n";
