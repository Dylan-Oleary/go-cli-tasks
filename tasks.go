package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type task struct {
    ID string `json:"id"`
    Description string `json:"description"`
    // TODO: Enums in Go?
    Status string `json:"status"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
type tasks []task

func (t tasks) add(value string) {
    var id string

    if(len(t) == 0) {
        id = "1"
    } else {
        last, err := strconv.Atoi(t[len(t) -1].ID)

        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }

        id = strconv.Itoa(last + 1)
    }

    taskToAdd := task{
        ID: id,
        Description: value,
        Status: "todo",
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
    }

    t = append(t, taskToAdd)

    writeTasksToFile(t)

    fmt.Printf("Task added successfully (ID: %s)", taskToAdd.ID)
}

func (t tasks) delete(id string) {
    for index, task := range t {
        if(task.ID == id) {
            t = append(t[:index], t[index + 1:]...)
            writeTasksToFile(t)

            fmt.Printf("Task (ID: %s) successfully deleted", id)
            return
        }
    }

    fmt.Printf("Task (ID: %s) does not exist", id)
    return
}

func (t tasks) update(id string, value string) {
    for i := range t {
        if(t[i].ID == id) {
            pointer := &t[i]

            pointer.Description = value
            pointer.UpdatedAt =  time.Now().UTC()

            fmt.Printf("Task (ID: %s) updated successfully)", id)
            writeTasksToFile(t)

            return
        }
    }

    fmt.Printf("Task (ID: %s) not found)", id)
}

func (t tasks) markDone(id string) {
    for i:= range t {
        if(t[i].ID == id) {
            pointer := &t[i]

            pointer.Status = "done"

            fmt.Printf("Task (ID: %s) marked as done", id)
            writeTasksToFile(t)

            return
        }
    }

    fmt.Printf("Task (ID: %s) does not exist", id)
    return
}

func (t tasks) markInProgress(id string) {
    for i:= range t {
        if(t[i].ID == id) {
            pointer := &t[i]

            pointer.Status = "in-progress"

            fmt.Printf("Task (ID: %s) marked as in-progress", id)
            writeTasksToFile(t)

            return
        }
    }

    fmt.Printf("Task (ID: %s) does not exist", id)
    return
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

func processCommand(c command, t tasks) error {
    action := c.action
    args := c.arguments

    switch action {
        case "add":
            if len(args) == 0 {
                return errors.New("Failed to add task due to missing task value")
            }

            t.add(args[0])
        case "delete":
            if len(args) == 0 {
                return errors.New("Failed to delete task due to missing task ID")
            }

            t.delete(args[0])
        case "list":
            if len(args) == 0 {
                t.list(args[0])
            }

            t.list()
        case "mark-done":
            if len(args) == 0 {
                return errors.New("Failed to mark task as 'done' due to missing task ID")
            }

            t.markDone(args[0])
        case "mark-in-progress":
            if len(args) == 0 {
                return errors.New("Failed to mark task as 'in-progress' due to missing task ID")
            }

            t.markInProgress(args[0])
        case "update":
            if len(args) < 2 {
                return errors.New("Failed to update task due to missing values")
            }

            t.update(args[0], args[1])
        default:
            return errors.New("Invalid action passed")
    }

    return nil
}

func printTasks(t tasks) {
    fmt.Println("Tasks")
    fmt.Println("-----")
    fmt.Println("")

    for _, task := range t {
        if(task.Status == "done"){
            fmt.Println("[X]", task.Description)
        } else {
            fmt.Println("[ ]", task.Description)
        }

        fmt.Println("    |-- ID:", task.ID)
        fmt.Println("    |-- Status:", task.Status)
        fmt.Println("    |-- Created At:", task.CreatedAt.Local().Format(time.RFC1123))
        fmt.Println("    |-- Updated At:", task.UpdatedAt.Local().Format(time.RFC1123))
        fmt.Println()
    }
}