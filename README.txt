# Wine Shop Inventory Management

This Go program manages the inventory of a wine shop, ensuring that 
the available wines are always sorted by ratings from wine experts, 
with the best-rated wine at the top. The program allows for the 
addition of new wines to the inventory and maintains the sorted list.

## Description

The Wine Shop Inventory Management program is designed to handle the 
dynamic inventory of a wine shop. It keeps the list of wines sorted 
by their ratings and allows for the addition of new wines with 
details such as (not in this order) country of origin, vineyard designation, 
winery, wine type, vintage, rating, and price per bottle.

## Features

- Read data from a CSV file
- Set the first row of data or header for the wine field(s)
- Add new wine(s) to the inventory
- Trim unnecessary spaces
- Sort wine(s) in descending order based on the ratings
- Append the entry date to each wine record in the wine inventory
- Output the list of wines using (';') as delimiter

## Installation

To run this program, you need to have Go installed on your system. Follow 
the steps below to set up and run the project:

 **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/wine-shop.git
   cd wine-shop


## Usage

1. Place your CSV file in the same directory as the program. Ensure the CSV file uses semicolons (`;`) as delimiters.

2. Run the program with the CSV file name as a command-line argument:

    ```sh
    go run main.go data.csv
    ```

3. The program will read the CSV file, process the wine record(s), 
add the entry date, and print the sorted wines inventory to the console.

## Example CSV File

Country of Origin,Vineyard designation,Winery,Wine type,Vintage,Rating,Price per bottle
 Chile,Valle Central,Baron De Rothschild,Chardonnay,2019,Parker: 94,1290
 Germany,Ihringer Winklerberg,Dr. Heger,Pinot Noir,2014,Robinson: 19,2390


## Code Overview

### Functions

- `trimSpaces(slice []string)`: Trims leading and trailing spaces from each element in the slice.
- `checkRowSize(row []string, expected_size int) bool`: Checks if each row has the expected number of fields (7).
- `searchRatingField(wine []string) (string, error)`: Searches for the field containing the rating value.
- `parseRating(ratingStr string) (float64, error)`: Parses the rating value from the string based on the rating scale.

### Main Program Flow

1. Opens the CSV file specified as a command-line argument.
2. Reads the header and data rows.
3. Trims spaces from the header and checks the row size.
4. Parses ratings and stores them in a map.
5. Appends the entry date(current date) to each record.
6. Sorts records based on ratings in descending order (Best Wine).
7. Prints the sorted wine list (inventory).