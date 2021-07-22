package main

import (
	"encoding/json"
	"fmt"
)

type FrequencyMap map[string]int

func getFrequency(str string) FrequencyMap {
	frequency := FrequencyMap{}
	for index := 0 ; index < len(str) ; index ++ {
		frequency[string( str[index] )] += 1
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
	strings := []string{"quick","brown","fox","lazy","dog"}
	frequencies := ParallelFrequency(strings)
	output , _ := json.MarshalIndent(frequencies, "", " ")
	fmt.Println(string(output))

}