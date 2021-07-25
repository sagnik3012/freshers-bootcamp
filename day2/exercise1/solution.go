package main

import (
	"encoding/json"
	"fmt"
)

type CharCount map[string]int

func getFreq(animal string) CharCount {
	freq := CharCount{}
	for idx := 0; idx < len(animal); idx++ {
		freq[string(animal[idx])] += 1
	}
	return freq
}

func ParaFreq(inputs []string) CharCount {
	output := CharCount{}
	nosOfStr := len(inputs)
	results := make(chan CharCount, nosOfStr)

	for _, currStr := range inputs {
		go func(s string) {
			results <- getFreq(s)
		}(currStr)
	}

	for idx := 0; idx < nosOfStr; idx++ {
		for char, freq := range <-results {
			output[char] += freq
		}
	}
	return output
}

func main() {

	animals := []string{"quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog"}

	frequencies := ParaFreq(animals)
	output, err := json.MarshalIndent(frequencies, "", " ")
	if err != nil {
		fmt.Println(" Error in converting to JSON \n :", err)
	} else {
		fmt.Println(string(output))
	}
}
