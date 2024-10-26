package main

func getActionsMap() map[string]string {
    return map[string]string{
        "add": "add",
        "delete": "delete",
        "list": "list",
        "mark-done": "mark-done",
        "mark-in-progress": "mark-in-progress",
        "update": "update",
    }
}

func getListActionsMap() map[string]string {
    return map[string]string{
        "done": "done",
        "in-progress": "in-progress",
        "todo": "todo",
    }
}