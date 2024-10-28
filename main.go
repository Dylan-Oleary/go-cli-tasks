package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
    t := readTasksFromFile()

    // Read CLI positional arguments
    args, err := getInputArgs()

    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    command := args[0]
    data := args[1:]

    switch command {
        case "add":
            if len(data) > 0 {
                t.add(data[0])
            } else {
                fmt.Println("Expected a value when attempting to add a value")
                os.Exit(1)
            }
        case "list":
            if len(data) > 0 {
                t.list(data[0])
            } else {
                t.list()
            }
        case "update":
            if len(data) == 2 {
                t.update(data[0], data[1])
            } else {
                fmt.Println("Invalid arguments passed when updating task")
                os.Exit(1)
            }
        default:
            fmt.Println("Unexpected action passed")
            os.Exit(1)
    }
}

func getInputArgs() ([]string, error) {
    args := os.Args

    if len(args) > 4 {
        return nil, errors.New("Too many arguments passed")
    }

    actionMap := getActionsMap()
    _, ok := actionMap[args[1]]

    if !ok {
        return nil, errors.New("Invalid action passed " + args[1])
    }

    return os.Args[1:], nil
}