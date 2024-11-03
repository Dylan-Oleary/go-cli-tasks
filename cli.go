package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
)

type command struct {
    // TODO ENUMS
    action string
    arguments []string
}

func getCommandValues() command {
    args := os.Args

    if len(args) <= 1 {
        printHelp()

        os.Exit(0)
    }

    action := args[1]
    actions := map[string]string {
        "add": "add",
        "delete": "delete",
        "list": "list",
        "mark-done": "mark-done",
        "mark-in-progress": "mark-in-progress",
        "update": "update",
    }
    _, ok := actions[action]

    if !ok {
        fmt.Println("Error: Invalid action passed –", action)
        fmt.Println("")
        fmt.Printf("Pass one of (%s)", slices.Sorted(maps.Keys(actions)))

        os.Exit(1)
    }

    arguments := args[2:]

    return command{action, arguments}
}


func printHelp() {
    fmt.Println("Usage:")
    fmt.Println("  task-cli <action> [arguments]")
    fmt.Println("")
    fmt.Println("Available actions:")
    fmt.Println("  add [task]  Add a new task -- task-cli add \"Eat pizza\"")
    fmt.Println("  delete [id]  Delete an existing task -- task-cli delete 1")
    fmt.Println("  list [done|in-progress]  List tasks")
    fmt.Println("  mark-done [id]  Mark an existing task as done -- task-cli mark-done 2")
    fmt.Println("  mark-in-progress [id]  Mark an existing task as in-progress -- task-cli mark-in-progress 3")
    fmt.Println("  update [id] [task]  Update an existing task -- task-cli update 1 \"Eat chocolate\"")
    fmt.Println("")
}