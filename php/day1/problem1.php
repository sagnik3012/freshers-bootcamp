<?php


$input = [2, 3, [4, 5], [6, 7], 8];
$output = [];

foreach ($input as $key => $value) {

    if (is_array($value)) {
        $output = array_merge($output, $value);
    } else {
        array_push($output, $value);
    }

}

echo "Output \n";
print_r($output);




