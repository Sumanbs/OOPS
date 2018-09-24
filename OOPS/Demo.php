<?php
class suman{
public $name="uman";
public $class = 10;
public $ne="jkhfdj";
}
$ref[2] = new suman();
$data = file_get_contents('demo.json');
file_put_contents('demo.json', json_encode($ref));
$data = file_get_contents('demo.json');
$json_arr = json_decode($data,true);
