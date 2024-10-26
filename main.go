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

// # Adding a new task
// task-cli add "Buy groceries"
// # Output: Task added successfully (ID: 1)

// # Updating and deleting tasks
// task-cli update 1 "Buy groceries and cook dinner"
// task-cli delete 1

// # Marking a task as in progress or done
// task-cli mark-in-progress 1
// task-cli mark-done 1