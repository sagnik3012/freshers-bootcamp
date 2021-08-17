<?php

$snakeCase = ["snake_case" , "camel_case"];
$camelCase = [];

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

foreach ( $snakeCase as $name){
    $camel = snakeToCamel($name);
    array_push($camelCase, $camel);
}

echo "Input : ";
print_r($snakeCase);
echo "Output : \n";
print_r($camelCase);
