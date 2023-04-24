package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

type Txn struct {
    ID          uint
    Date        string
    Transaction float64
}

func readCSV(filename string) ([]Txn, error) {
    // Open the CSV file
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Create a new CSV reader
    reader := csv.NewReader(file)

    // Read all the records
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    // Create an array to store the Txn objects
    txns := make([]Txn, 0)

    // Loop through the records (excluding the header row)
    for i, record := range records {
        if i == 0 {
            continue // skip header row
        }
        // Create a new Txn object and append it to the array
        txn, err := createTxn(record)
        if err != nil {
            return nil, err
        }
        txns = append(txns, txn)
    }

    return txns, nil
}

func createTxn(line []string) (Txn, error) {
    // Parse the values
    idStr := line[0]
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        return Txn{}, err
    }

    date := line[1]
    transactionStr := line[2]

    // Parse the transaction amount
    transaction, err := strconv.ParseFloat(transactionStr, 64)
    if err != nil {
        return Txn{}, err
    }

    // Create a new Txn object and return it
    return Txn{
        ID:          uint(id),
        Date:        date,
        Transaction: transaction,
    }, nil
}

func main() {
    txns, err := readCSV("../txns.csv")
    if err != nil {
        fmt.Println("Error reading CSV file:", err)
        return
    }

    // Print the Txn objects
    for _, txn := range txns {
        fmt.Printf("ID: %d, Date: %s, Transaction: %.2f\n", txn.ID, txn.Date, txn.Transaction)
    }
}
