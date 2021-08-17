
<?php

class Cricketers{

    public function __construct($inputJson){
        $this->data = json_decode($inputJson,true);

    }

    // prints all names, ages and cities
    public function getNamesAgesAndCities(){
        $name = [] ; $age = [] ; $city = [];
        foreach ($this->data["players"] as $values){
            array_push($name,$values["name"]);
            array_push($age,$values["age"]);
            array_push($city,$values["address"]["city"]);
        }
        echo "Names:\n";
        print_r($name);
        echo "Ages:\n";
        print_r($age);
        echo "Cities:\n";
        print_r($city);
    }

    // gets unique names
    public function getUniqueNames(){
        $names = [];
        foreach( $this->data["players"] as $cricketer ){
            array_push($names , $cricketer["name"]);
        }
        $names = array_map( 'strtolower', $names );
        $uniqueNames = array_unique($names);
        $uniqueNames = array_map( 'ucfirst', $uniqueNames );
        return $uniqueNames;
    }

    // gets names of people with max age
    public function getPeopleWithMaxAge(){
        $result = [];
        $maxAge = 0;
        foreach ($this->data["players"] as $cricketer){

            if ($cricketer["age"]>$maxAge){
                $maxAge = $cricketer["age"];
            }
        }
        foreach($this->data["players"] as $cricketer){

            if ( $cricketer["age"] == $maxAge){
                array_push($result,$cricketer["name"]);
            }
        }
        return $result;
    }

}

$CricketersJSON = "{\"players\":[{\"name\":\"Ganguly\",\"age\":45,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Dravid\",\"age\":45,\"address\":{\"city\":\"Bangalore\"}},{\"name\":\"Dhoni\",\"age\":37,\"address\":{\"city\":\"Ranchi\"}},{\"name\":\"Virat\",\"age\":35,\"address\":{\"city\":\"Delhi\"}},{\"name\":\"jadeja\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Jadeja\",\"age\":35,\"address\":{\"city\":\"Mumbai\"}}]}";
$cricketers = new Cricketers($CricketersJSON);

// part 1 solution
echo "Solution to part 1 :\n";
$cricketers->getNamesAgesAndCities();
// part 2 solution
echo "Solution to part 2 :\n";
echo "Unique Cricketer Names:\n";
print_r($cricketers->getUniqueNames());
// part 3 solution
echo "Solution to part 3 :\n";
print "Cricketers with maximum age :\n";
print_r($cricketers->getPeopleWithMaxAge());