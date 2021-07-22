package main

import (
	"encoding/json"
	"fmt"
)

type FrequencyMap map[string]int

func getFrequency(animal string) FrequencyMap {
	frequency := FrequencyMap{}
	for index := 0 ; index < len(animal) ; index ++ {
		frequency[string( animal[index] )] += 1
	}
	return frequency
}

func ParallelFrequency(inputStrings []string) FrequencyMap {
	output := FrequencyMap{}
	numberofStrings := len(inputStrings)
	results := make(chan FrequencyMap, numberofStrings)

	for _, currentString := range inputStrings{
		go func(s string) {
			results <- getFrequency(s)
		}(currentString)
	}

	for index := 0; index < numberofStrings ; index++ {
		for character , frequency:= range <-results {
			output[ character ] += frequency
		}
	}
	return output
}

func main(){

	animals := []string{"quick","brown","fox","lazy","dog","happy","cat"}

	frequencies := ParallelFrequency(animals)
	output , err := json.MarshalIndent(frequencies, "", " ")
	if err != nil {
		fmt.Println(" Error in converting to JSON \n ",err)
	} else{
		fmt.Println(string(output))
	}
}