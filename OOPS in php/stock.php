<?php
/**
 * Desc -> Write a program to read in Stock Names, Number of Share, Share Price.
 *Print a Stock Report with total value of each Stock and the total value of Stock.
 */
class Stock
{
    /**
     * @var string
     * @var int
     * @var int
     * @var int
     */
    public $name;
    public $stock_price;
    public $number_of_shares;
   // public $total_price;

    public function __construct($name,$number_of_shares,$stock_price)
    {
        $this->name = $name;
        $this->stock_price = $stock_price;
        $this->number_of_shares = $number_of_shares;
        //$this->total_price = $number_of_shares * $stock_price;
    }
    public function add_to_json()
    {
        $data = file_get_contents('demo.json');
        $json_arr = json_decode($data,true);
        $json_arr[] = array('name'=>$this->name,'stock_price'=>$this->stock_price,'number_of_shares'=>$this->number_of_shares);
        file_put_contents('demo.json', json_encode($json_arr));
    }
}
?>