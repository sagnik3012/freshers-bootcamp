<?php

$input = ["snake_case" , "camel_case" , "snake_snake_camel_camel"];

function snakeToCamel( $snake ){

    $camel = $snake;
    for( $i  = 0 ; $i < strlen($camel) -1 ; $i ++ ){

        if ( $camel[$i] == '_'){
            $camel[$i+1] = strtoupper($camel[$i+1]);
        }
    }
    $camel = str_replace("_" , "",$camel);
    return $camel;

}
$output = [];

foreach ( $input as $name){
    $camel = snakeToCamel($name);
    array_push($output, $camel);
}

print_r($output);

