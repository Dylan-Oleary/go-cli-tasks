package main

import (
	"fmt"
	"os"
)

func main() {
    command := getCommandValues()
    tasks := readTasksFromFile()

    err := processCommand(command, tasks)

    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(0)
    }
}