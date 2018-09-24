<?php
/**
 * Read the inventory details from JSON file and display the price.
 */
$json = file_get_contents("inventory.json");
/**
 * Decode the json contents to array of objects.
 */
$json_data = json_decode($json);

//Calculate total Price
$rice_price = $json_data->Rice->price * $json_data->Rice->weight;
$wheat_price = $json_data->Wheat->price * $json_data->Wheat->weight;
$pulses_price = $json_data->pulses->price * $json_data->pulses->weight;
//display the details here
echo "Rice Details : \n";
echo "Name := " . $json_data->Rice->name . "   TotalPrice := " . $rice_price;
echo "\nWheat Details :\n ";
echo "Name := " . $json_data->Wheat->name . "   TotalPrice := " . $wheat_price;
echo "\nPulses Details : \n";
echo "Name := " . $json_data->pulses->name . "  TotalPrice := " . $pulses_price;
echo "\n"
?>
