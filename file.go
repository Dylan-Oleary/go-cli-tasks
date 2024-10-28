package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const FILE_NAME = "data.json"

func readTasksFromFile() tasks {
    _, checkIfFileExistsError := os.Stat(FILE_NAME)

    if errors.Is(checkIfFileExistsError, os.ErrNotExist) {
        fmt.Println("Unable to find file:", FILE_NAME, "– creating new file.")

        data := []byte("[]")
        os.WriteFile(FILE_NAME, data, 0666)
    }

    bytes, readFileError := os.ReadFile(FILE_NAME)
    
    if readFileError != nil {
        fmt.Println("Failed to read file:", FILE_NAME)
        fmt.Println("Error:", readFileError)
        os.Exit(1)
    }

    tasks := []task{}
    jsonParseError := json.Unmarshal(bytes, &tasks)

    if jsonParseError != nil {
        fmt.Println("Failed to parse file:", FILE_NAME)
        fmt.Println("Error:", jsonParseError)
        os.Exit(1)
    }

    return tasks
}

func writeTasksToFile(t tasks) {
    data, byteConversionError := json.Marshal(t)

    if byteConversionError != nil {
        fmt.Println("Failed to convert task list to a byte array")
        fmt.Println("Error:", byteConversionError)
        os.Exit(1)
    }

    writeFileError := os.WriteFile(FILE_NAME, []byte(data), 0666)

    if writeFileError != nil {
        fmt.Println("Failed to write new tasks to file:", FILE_NAME)
        fmt.Println("Error:", writeFileError)
        os.Exit(1)
    }
}