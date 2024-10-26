package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func readTasksFromFile(filePath string) tasks {
    _, checkIfFileExistsError := os.Stat(filePath)

    if errors.Is(checkIfFileExistsError, os.ErrNotExist) {
        fmt.Println("Unable to find file:", filePath, "– creating new file.")

        data := []byte("[]")
        os.WriteFile(filePath, data, 0666)
    }

    bytes, readFileError := os.ReadFile(filePath)
    
    if readFileError != nil {
        fmt.Println("Failed to read file:", filePath)
        fmt.Println("Error:", readFileError)
        os.Exit(1)
    }

    tasks := []task{}
    jsonParseError := json.Unmarshal(bytes, &tasks)

    if jsonParseError != nil {
        fmt.Println("Failed to parse file:", filePath)
        fmt.Println("Error:", jsonParseError)
        os.Exit(1)
    }

    return tasks
}

func writeTasksToFile(filePath string, t tasks) {
    data, byteConversionError := json.Marshal(t)

    if byteConversionError != nil {
        fmt.Println("Failed to convert task list to a byte array")
        fmt.Println("Error:", byteConversionError)
        os.Exit(1)
    }

    writeFileError := os.WriteFile(filePath, []byte(data), 0666)

    if writeFileError != nil {
        fmt.Println("Failed to write new tasks to file:", filePath)
        fmt.Println("Error:", writeFileError)
        os.Exit(1)
    }
}