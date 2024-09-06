package main

import (
	"bytes"
	"math/big"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoHandler(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	fileWriter, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("could not create form file: %v", err)
	}

	// Write CSV content to the file part
	csvContent := "1,2,3\n4,5,6\n7,8,9\n"
	if _, err := fileWriter.Write([]byte(csvContent)); err != nil {
		t.Fatalf("could not write to form file: %v", err)
	}

	// Close the writer to set the correct Content-Type boundary
	writer.Close()

	// Create a new request with the body
	req := httptest.NewRequest(http.MethodPost, "/echo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(echoHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestInvertHandler(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	fileWriter, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("could not create form file: %v", err)
	}

	// Write CSV content to the file part
	csvContent := "1,2,3\n4,5,6\n7,8,9\n"
	if _, err := fileWriter.Write([]byte(csvContent)); err != nil {
		t.Fatalf("could not write to form file: %v", err)
	}

	// Close the writer to set the correct Content-Type boundary
	writer.Close()

	// Create a new request with the body
	req := httptest.NewRequest(http.MethodPost, "/invert", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(invertHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "1,4,7\n2,5,8\n3,6,9\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestInvertHandlerIncompleteMatrix(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	fileWriter, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("could not create form file: %v", err)
	}

	// Write CSV content to the file part
	csvContent := "1,2,3\n4,5,6\n"
	if _, err := fileWriter.Write([]byte(csvContent)); err != nil {
		t.Fatalf("could not write to form file: %v", err)
	}

	// Close the writer to set the correct Content-Type boundary
	writer.Close()

	// Create a new request with the body
	req := httptest.NewRequest(http.MethodPost, "/invert", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(invertHandler)
	handler.ServeHTTP(rr, req)

	// Check the response body
	expected := "incomplete matrix"
	assert.Contains(t, rr.Body.String(), expected)
}
func TestFlattenHandler(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	fileWriter, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("could not create form file: %v", err)
	}

	// Write CSV content to the file part
	csvContent := "1,2,3\n4,5,6\n7,8,9\n"
	if _, err := fileWriter.Write([]byte(csvContent)); err != nil {
		t.Fatalf("could not write to form file: %v", err)
	}

	// Close the writer to set the correct Content-Type boundary
	writer.Close()

	// Create a new request with the body
	req := httptest.NewRequest(http.MethodPost, "/flatten", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(flattenHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "1,2,3,4,5,6,7,8,9\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestFlattenHandlerIncompleteMatrix(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	fileWriter, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("could not create form file: %v", err)
	}

	// Write CSV content to the file part
	csvContent := "1,2,3\n"
	if _, err := fileWriter.Write([]byte(csvContent)); err != nil {
		t.Fatalf("could not write to form file: %v", err)
	}

	// Close the writer to set the correct Content-Type boundary
	writer.Close()

	// Create a new request with the body
	req := httptest.NewRequest(http.MethodPost, "/invert", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(invertHandler)
	handler.ServeHTTP(rr, req)

	// Check the response body
	expected := "csv with incomplete matrix"
	assert.Contains(t, rr.Body.String(), expected)
}
func TestSumHandler(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	fileWriter, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("could not create form file: %v", err)
	}

	// Write CSV content to the file part
	csvContent := "1,2,3\n4,5,6\n7,8,9\n"
	if _, err := fileWriter.Write([]byte(csvContent)); err != nil {
		t.Fatalf("could not write to form file: %v", err)
	}

	// Close the writer to set the correct Content-Type boundary
	writer.Close()

	// Create a new request with the body
	req := httptest.NewRequest(http.MethodPost, "/sum", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(sumHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "45\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestMultiplyHandler(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	fileWriter, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("could not create form file: %v", err)
	}

	// Write CSV content to the file part
	csvContent := "1,2,3\n4,5,6\n7,8,9\n"
	if _, err := fileWriter.Write([]byte(csvContent)); err != nil {
		t.Fatalf("could not write to form file: %v", err)
	}

	// Close the writer to set the correct Content-Type boundary
	writer.Close()

	// Create a new request with the body
	req := httptest.NewRequest(http.MethodPost, "/multiply", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(multiplyHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "362880\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestSumRow(t *testing.T) {
	input := []string{"1", "2", "3"}
	expected := big.NewInt(6)
	result, err := sumRow(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assert.Equal(t, expected, result)

}

func TestMultiplyRow(t *testing.T) {
	input := []string{"4", "5", "6"}
	expected := big.NewInt(120)
	result, err := multiplyRow(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assert.Equal(t, expected, result)

}

func TestMatrixToString(t *testing.T) {
	input := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	result := matrixToString(input)
	assert.Equal(t, expected, result)

}
