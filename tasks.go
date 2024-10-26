package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type task struct {
    ID string `json:"id"`
    Description string `json:"description"`
    // TODO: Enums in Go?
    Status string `json:"status"`
    // TODO: Datetime in Go?
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}
type tasks []task

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

func (t tasks) list(filter ...string) {
    if(filter == nil) {
        printTasks(t)
        return
    }

    listFilters := map[string]string{
        "done": "done",
        "in-progress": "in-progress",
        "todo": "todo",
    }

    validFilter, ok := listFilters[filter[0]]

    if !ok {
        fmt.Println("Invalid list filter passed:", filter[0])
        os.Exit(1)
    }

    filteredTasks := tasks{}

    for _, task := range t {
        if(task.Status == validFilter) {
            filteredTasks = append(filteredTasks, task)
        }
    }

    printTasks(filteredTasks)
    return
}

func printTasks(t tasks) {
    fmt.Println("Tasks")
    fmt.Println()

    for _, task := range t {
        fmt.Println("ID:", task.ID)
        fmt.Println("Description:", task.Description)
        fmt.Println("Status:", task.Status)
        fmt.Println("Created At:", task.CreatedAt)
        fmt.Println("Updated At:", task.UpdatedAt)
        fmt.Println()
    }
}