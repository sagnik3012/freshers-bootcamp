

<?php

$phoneNumber = 9876543210;

function maskPhoneNumber( $phoneNumber){

    $number = strval($phoneNumber);
    return substr($number , 0 , 2)
        . str_repeat("*",6).substr($number,-2);

}
echo maskPhoneNumber($phoneNumber);



