<?php
class Address1
{
    /**
     * @var string FirstName
     * @var string Lastname
     * @var array addresss
     */
     
    public $FirstName;
    public $LastName;
    public $address = array();
    /**
     * constructor with parameters
     */
    public function __construct($FirstName, $LastName, $state, $city, $zipcode, $phone_Number)
    {
        $this->FirstName = $FirstName;
        $this->LastName = $LastName;
        $this->address = array("State" => $state, "City" => $city,
            "Zipcode" => $zipcode,
            "PH-NO" => $phone_Number);
    }
}