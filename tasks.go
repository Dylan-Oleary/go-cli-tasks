package main

import (
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

// task-cli delete 1
func (t tasks) delete(id string) {

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

// task-cli mark-in-progress 1
// task-cli mark-done 1
func (t tasks) updateStatus(status string) {

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