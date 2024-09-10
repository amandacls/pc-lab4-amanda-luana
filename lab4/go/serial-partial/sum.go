package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filePath string) (int, error) {
	data, err := readFile(filePath)
	if err != nil {
		return 0, err
	}

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	return _sum, nil
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)

	for _, path := range os.Args[1:] {
		_sum, err := sum(path)

		if err != nil {
			continue
		}

		totalSum += int64(_sum)

		sums[_sum] = append(sums[_sum], path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}

	//fileFingerprints := make(map[string]int)

	for i := 0; i < len(os.Args); i++ {
		for j := i + 1; j < len(os.Args); j++ {
			file1 := os.Args[i]
			file2 := os.Args[j]

			// Sem funcionar pela ausencia do código do somador de chunks

			// fingerprint1 := sums[file1]
			// fingerprint2 := sums[file2]
			// similarityScore := similarity(fingerprint1, fingerprint2)
			fmt.Printf("Similarity between %s and %s: %.2f%%\n", file1, file2)
			//, similarityScore*100)
		}
	}
}

func similarity(base, target []int64) float64 {
	counter := 0
	targetMap := make(map[int64]bool)

	for _, value := range target {
		targetMap[value] = true
	}

	for _, value := range base {
		if _, exists := targetMap[value]; exists {
			counter++
			delete(targetMap, value)
		}
	}

	return float64(counter) / float64(len(base))
}
