package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to fetch csv file: %v", err)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Printf("failed to read csv file: %v", err)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(w, response)
}

func invertHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to fetch file: %v", err)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var inverted [][]string
	var rowIndex int
	var rowCount, size int
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				if rowCount == 0 {
					log.Printf("csv is empty: %v", err)
					http.Error(w, "csv is empty", http.StatusBadRequest)
					return
				}

				if rowCount < size {
					diff_str := fmt.Sprintf("csv with incomplete matrix. Expected %dx%d, Got: %dx%d", size, size, rowCount, size)
					log.Printf("%s: %v", diff_str, err)
					http.Error(w, diff_str, http.StatusBadRequest)
					return
				}
				break
			}

			log.Printf("error while reading file: %v", err)
			http.Error(w, fmt.Sprintf("error reading file: %s", err.Error()), http.StatusBadRequest)
			return
		}
		rowCount += 1
		// Update size if first row
		if rowCount == 1 {
			size = len(row)
		}

		if rowCount > size {
			diff_str := fmt.Sprintf("csv with incomplete matrix. Expected %dx%d, Got: %dx%d", size, size, rowCount, size)
			log.Printf("%s: %v", diff_str, err)
			http.Error(w, diff_str, http.StatusBadRequest)
			return
		}

		for colIndex := range row {
			if len(inverted) <= colIndex {
				inverted = append(inverted, make([]string, len(row)))
			}

			inverted[colIndex][rowIndex] = row[colIndex]
		}

		rowIndex++
	}

	result := matrixToString(inverted)

	fmt.Fprint(w, result)

}

func flattenHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to fetch file: %v", err)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var response string
	var rowCount, size int
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				if rowCount == 0 {
					log.Printf("csv is empty: %v", err)
					http.Error(w, "csv is empty", http.StatusBadRequest)
					return
				}

				if rowCount < size {
					diff_str := fmt.Sprintf("csv with incomplete matrix. Expected %dx%d, Got: %dx%d", size, size, rowCount, size)
					log.Printf("%s: %v", diff_str, err)
					http.Error(w, diff_str, http.StatusBadRequest)
					return
				}

				break
			}
			log.Printf("error while reading file: %v", err)
			http.Error(w, fmt.Sprintf("error reading file: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		rowCount += 1
		// Update size if first row
		if rowCount == 1 {
			size = len(row)
		}

		if rowCount > size {
			diff_str := fmt.Sprintf("csv with incomplete matrix. Expected %dx%d, Got: %dx%d", size, size, rowCount, size)
			log.Printf("%s: %v", diff_str, err)
			http.Error(w, diff_str, http.StatusBadRequest)
			return
		}
		response += strings.Join(row, ",") + ","
	}
	if len(response) > 0 {
		response = response[:len(response)-1]
	}

	fmt.Fprintln(w, response)

}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to fetch file: %v", err)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	totalSum := big.NewInt(0)
	var rowCount, size int
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				if rowCount == 0 {
					log.Printf("csv is empty: %v", err)
					http.Error(w, "csv is empty", http.StatusBadRequest)
					return
				}

				if rowCount < size {
					diff_str := fmt.Sprintf("csv with incomplete matrix. Expected %dx%d, Got: %dx%d", size, size, rowCount, size)
					log.Printf("%s: %v", diff_str, err)
					http.Error(w, diff_str, http.StatusBadRequest)
					return
				}

				break
			}
			log.Printf("error while reading file: %v", err)
			http.Error(w, fmt.Sprintf("error reading file: %s", err.Error()), http.StatusBadRequest)
			return
		}
		rowCount += 1
		// Update size if first row
		if rowCount == 1 {
			size = len(row)
		}

		if rowCount > size {
			diff_str := fmt.Sprintf("csv with incomplete matrix. Expected %dx%d, Got: %dx%d", size, size, rowCount, size)
			log.Printf("%s: %v", diff_str, err)
			http.Error(w, diff_str, http.StatusBadRequest)
			return
		}
		sum, err := sumRow(row)
		if err != nil {
			log.Printf("error summing row : %v", err)
			http.Error(w, fmt.Sprintf("error summing row: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		totalSum.Add(totalSum, sum)
	}

	fmt.Fprintln(w, totalSum)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to fetch file: %v", err)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	totalProduct := big.NewInt(1)
	var rowCount, size int
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				if rowCount == 0 {
					log.Printf("csv is empty: %v", err)
					http.Error(w, "csv is empty", http.StatusBadRequest)
					return
				}

				if rowCount < size {
					diff_str := fmt.Sprintf("csv with incomplete matrix. Expected %dx%d, Got: %dx%d", size, size, rowCount, size)
					log.Printf("%s: %v", diff_str, err)
					http.Error(w, diff_str, http.StatusBadRequest)
					return
				}

				break
			}
			log.Printf("error while reading file: %v", err)
			http.Error(w, fmt.Sprintf("error reading file: %s", err.Error()), http.StatusBadRequest)
			return
		}

		rowCount += 1
		// Update size if first row
		if rowCount == 1 {
			size = len(row)
		}

		if rowCount > size {
			diff_str := fmt.Sprintf("csv with incomplete matrix. Expected %dx%d, Got: %dx%d", size, size, rowCount, size)
			log.Printf("%s: %v", diff_str, err)
			http.Error(w, diff_str, http.StatusBadRequest)
			return
		}

		product, err := multiplyRow(row)
		if err != nil {
			log.Printf("error multiplying row : %v", err)
			http.Error(w, fmt.Sprintf("error multiplying row: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// If product is 0, early return
		if product == big.NewInt(0) {
			fmt.Fprintln(w, 0)
			return
		}

		totalProduct.Mul(totalProduct, product)
	}

	fmt.Fprintln(w, totalProduct)

}

func matrixToString(matrix [][]string) string {
	var result string
	for _, row := range matrix {
		result += strings.Join(row, ",") + "\n"
	}
	return result
}

func sumRow(row []string) (*big.Int, error) {
	sum := big.NewInt(0)
	for _, value := range row {
		// Convert string to integer
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("invalid number %s: %v", value, err)
			return nil, fmt.Errorf("invalid number %s: %w", value, err)
		}
		sum.Add(sum, big.NewInt(int64(num)))
	}
	return sum, nil
}

func multiplyRow(row []string) (*big.Int, error) {
	product := big.NewInt(1)
	for _, value := range row {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("invalid number %s: %w", value, err)
		}
		product.Mul(product, big.NewInt(int64(num)))
	}
	return product, nil
}
