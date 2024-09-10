package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Result struct {
	Sum int
	Path string
}

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

func sumRoutine(filePath string, out chan Result) {
	//filePath := <-in
	_sum, _ := sum(filePath)
	_result := Result{_sum, filePath}
	out <- _result
}
// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
	result := make(chan Result)
	for _, path := range os.Args[1:] {
		go sumRoutine(path, result)
	}

	for i:=0; i < len(os.Args[1:]); i++ {
		_result := <-result
		_sum := _result.Sum
		_filePath := _result.Path
		totalSum += int64(_sum)
		sums[_sum] = append(sums[_sum], _filePath)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}


