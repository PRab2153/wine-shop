package main

import (
    "log"
    "flag"
	"fmt"
	"encoding/csv"
    "os"
    "strings"
    "strconv"
	"time"
    "sort"
)

// Wine data fields' size
const WINE_SIZE = 7

// trimSpaces for unnecessary spaces
func trimSpaces(slice []string) {
	for i := range slice {
		slice[i] = strings.TrimSpace(slice[i])
	}
}

// Check if the row is exceeding or missing from the data or not
func checkRowSize(row []string, expected_size int) bool {
	return len(row) == expected_size
}

// findRating searches for a colon ':' in one of the fields of a row
func findRating(wine []string) (string, error){
    for _, field := range wine {
		if strings.Contains(field, ":") {

			return field, nil
		}
	}
	return "", fmt.Errorf("Rating field not found")
}

// parseRating parses the rating value from the string based on the rating scale
func parseRating(ratingStr string) (float64, error) {
	parts := strings.Split(ratingStr, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("Invalid rating format: %s", ratingStr)
	}

	scale := strings.TrimSpace(parts[0])
	value, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	
    if err != nil {
		return 0, fmt.Errorf("Failed to parse rating value: %w", err)
	}

	switch scale {
        //Use Parker as standard
	case "Parker":
		return value, nil
	case "Robinson":
		// Convert Robinson scale to Parker scale
		return value * 5, nil

	default:
		return 0, fmt.Errorf("Invalid rating scale: %s", scale)
	}
}


func main() {
    // Part 1 Read and Check the file
    // Define and parse command-line arguments
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatalf("Input file is required. Usage: go run main.go <input-file>")
	}
	inputFile := flag.Arg(0)

    // Open the CSV file
    file, err := os.Open(inputFile)
    if err != nil {
        log.Fatalf("Failed to open file: %s", err)
    }
    defer file.Close()

    // Create a new CSV reader
    reader := csv.NewReader(file)

    // Read the header
    header, err := reader.Read()
    if err != nil {
        log.Fatalf("Failed to read header: %s", err)
    }

    // Read the details of wine record(s)
    wines, err := reader.ReadAll()
    if err != nil {            
        log.Fatalf("Failed to read records: %s", err)
    }

    // Trim spaces from the header
	trimSpaces(header)

	// Check if the header meets the required size
	if !checkRowSize(header, WINE_SIZE) {
		log.Fatalf("Header size mismatch: expected %d columns, got %d", WINE_SIZE, len(header))
	}

    // Define a map to store ratings for sorting
	ratings := make(map[string]float64)
    // Get the current date and format it as a string
	entryDate := time.Now().Format("2006-01-02")

    // Trim spaces from each record of wine and check their size
	for i, wine := range wines {

        //Clear unnecessary spaces
		trimSpaces(wine)

        //Check if the row's size meets the requirement
		if !checkRowSize(wines[i], WINE_SIZE) {
			log.Fatalf("Record size mismatch: expected %d columns, got %d", WINE_SIZE, len(wines[i]))
		}

        //Search for rating field
        ratingField, err := findRating(wine)
		if err != nil {
			log.Fatalf("Failed to find rating field in record: %v", err)
		}

        //Parsing the rating field
        rating, err := parseRating(ratingField)
		if err != nil {
			log.Fatalf("Failed to parse rating for record: %s", err)
		}

        //Store the wine record with entry date into temporary list to compare rating between wines
		ratings[strings.Join(append(wine, entryDate), ";")] = rating
	}

    //Part 2 Sorting and Displaying
    // Define a sorted list of wine bottles
    var sortedWines []string

    // Append wine into the wine list in descending order (Best Wine)
	for wine := range ratings {
		sortedWines = append(sortedWines, wine)
	}
	sort.Slice(sortedWines, func(i, j int) bool {
		return ratings[sortedWines[i]] > ratings[sortedWines[j]]
	})

    //Format Header
    headerString := strings.Join(append(header, "Entry Date"), ";")
	
    // Display header and sorted wines list
    fmt.Println(headerString)
	for _, wine := range sortedWines {
        fmt.Println(wine)
	}
}