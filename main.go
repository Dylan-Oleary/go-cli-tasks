package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
    t := readTasksFromFile()

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
        case "delete":
            if len (data) > 0 {
                t.delete(data[0])
            } else {
                fmt.Println("Expected an id when deleting a task")
                os.Exit(1)
            }
        case "list":
            if len(data) > 0 {
                t.list(data[0])
            } else {
                t.list()
            }
        case "mark-done":
            if len (data) > 0 {
                t.markDone(data[0])
            } else {
                fmt.Println("Expected an id when marking task as done")
                os.Exit(1)
            }
        case "mark-in-progress":
            if len (data) > 0 {
                t.markInProgress(data[0])
            } else {
                fmt.Println("Expected an id when marking task as in-progress")
                os.Exit(1)
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

    if len(args) == 1 {
        printHelp()
        os.Exit(1)
    }

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

func printHelp() {
    fmt.Println("Usage:")
    fmt.Println("  task-cli <action> [arguments]")
    fmt.Println("")
    fmt.Println("Available actions:")
    fmt.Println("  add <task>  Add a new task -- task-cli add \"Eat pizza\"")
    fmt.Println("  delete <id>  Delete an existing task -- task-cli delete 1")
    fmt.Println("  list  List all tasks -- task-cli list")
    fmt.Println("    list done  List all tasks marked as \"done\" -- task-cli list done")
    fmt.Println("    list in-progress  List all tasks marked as \"in-progress\" -- task-cli list in-progess")
    fmt.Println("  mark-done <id>  Mark an existing task as done -- task-cli mark-done 2")
    fmt.Println("  mark-on-progress <id>  Mark an existing task as in-progress -- task-cli mark-in-progress 3")
    fmt.Println("  update <id> <task>  Update an existing task -- task-cli update 1 \"Eat chocolate\"")
    fmt.Println("")
}