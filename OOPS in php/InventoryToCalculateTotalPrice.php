<?php
class Inventory
{
    public function getInventory()
    {
        /**
         * get the contents from json file.
         */
        $data = file_get_contents('inventory.json');
        /**
         * Decode it to array
         */
        $json_data = json_decode($data);
        //calculate Price of each inventory
        $rice_price = $json_data->Rice->price * $json_data->Rice->weight;
        echo "\nRice price = " . $rice_price . "\n";
        $wheat_price = $json_data->Wheat->price * $json_data->Wheat->weight;
        echo "wheat_price = " . $wheat_price . "\n";
        $pulses_price = $json_data->pulses->price * $json_data->pulses->weight;
        echo "pulses_price = " . $pulses_price . "\n";

        //calculate total price
        $total_Inventory = $rice_price + $wheat_price + $pulses_price;
        echo "\nTotal Inventory amount = " . $total_Inventory . "\n";
    }
}
$ref = new Inventory();
$ref->getInventory();
