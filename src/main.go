package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// struct for the matrix file
type SparceMatrixFile struct {
	row  int
	col  int
	data map[[2]int]int
}

func SpaceMatrix(rows int, cols int) *SparceMatrixFile {
	return &SparceMatrixFile{
		row:  rows,
		col:  cols,
		data: make(map[[2]int]int),
	}
}

func LoadDataFromFile(filePath string) (*SparceMatrixFile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	// This will close the file incase the function is done running
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows, cols int
	values := make(map[[2]int]int)

	if scanner.Scan() {
		rows, err = strconv.Atoi(strings.Split(scanner.Text(), "=")[1])
		if err != nil {
			return nil, errors.New("Invalid row")
		}
	}

	if scanner.Scan() {
		cols, err = strconv.Atoi(strings.Split(scanner.Text(), "=")[1])
		if err != nil {
			return nil, errors.New("Invalid colum")
		}
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		line = strings.Trim(line, "()")
		parts := strings.Split(line, ",")

		// if the length of individual matrix not three throw error
		if len(parts) != 3 {
			return nil, errors.New("Wrong matrix format")
		}

		// get the row, col and val from the string text
		row, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		col, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		val, err3 := strconv.Atoi(strings.TrimSpace(parts[2]))

		// Incase of error from row, cols, val return error
		if err1 != nil || err2 != nil || err3 != nil {
			return nil, errors.New("Invalid matrix values")
		}

		values[[2]int{row, col}] = val
	}

	return &SparceMatrixFile{rows, cols, values}, nil
}

// Function to help in printing the results
func (m *SparceMatrixFile) Print() {
	for key, value := range m.data {
		fmt.Printf("(%d, %d,%d)\n", key[0], key[1], value)
	}
}

func addition(file1 *SparceMatrixFile, file2 *SparceMatrixFile) (*SparceMatrixFile, error) {
	// check if the matrices are compatible
	if file1.row != file2.row || file1.col != file2.col {
		return nil, errors.New("matrices have incompatible dimensions")
	}

	result := SpaceMatrix(file1.row, file1.col)

	for key, value := range file1.data {
		result.data[key] = value
	}

	for key, value := range file2.data {
		if existingValue, exists := result.data[key]; exists {
			result.data[key] = existingValue + value
		} else {
			result.data[key] = value
		}
		if result.data[key] == 0 {
			delete(result.data, key)
		}
	}
	return result, nil
}

func subtract(file1 *SparceMatrixFile, file2 *SparceMatrixFile) (*SparceMatrixFile, error) {
	// check if the matrices are compatible
	if file1.row != file2.row || file1.col != file2.col {
		return nil, errors.New("matrices have incompatible dimensions")
	}

	result := SpaceMatrix(file1.row, file1.col)

	for key, value := range file1.data {
		result.data[key] = value
	}

	for key, value := range file2.data {
		// if exists in both value subtract them
		if existingValue, exists := result.data[key]; exists {
			result.data[key] = existingValue - value
		} else {
			result.data[key] = value
		}
		if result.data[key] == 0 {
			delete(result.data, key)
		}
	}
	return result, nil
}

func main() {
	file1, err := LoadDataFromFile("./sample_inputs/easy_sample_03_1.txt")
	if err != nil {
		fmt.Println("Error loading file1:", err)
		return
	}

	file2, err := LoadDataFromFile("./sample_inputs/easy_sample_03_2.txt")
	if err != nil {
		fmt.Println("Error loading file2:", err)
		return
	}

	// Addition of the two files
	addition, err := addition(file1, file2)
	if err != nil {
		fmt.Println("Error adding:", err)
		return
	}
	addition.Print()

	// Subtraction of the two files
	subtraction, err := subtract(file1, file2)
	if err != nil {
		fmt.Println("Error subtracting:", err)
		return
	}
	subtraction.Print()

	fmt.Println("Successfull Operations done!")
}
