package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
    t := readTasksFromFile("tasks.json")

    // Read CLI positional arguments
    _, err := getInputArgs()

    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // TODO
    // Action Handler
    t.list()
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

    return os.Args, nil
}